package route

import (
	"github.com/gin-gonic/gin"

	"github.com/gotokatsuya/lboard/controller"
)

// Init ...
func Init() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("view/**/*")
	router.Static("/public", "public")
	renderGroup := router.Group("/render")
	{
		renderGroup.GET("/", controller.RenderIndex)
		renderGroup.GET("/posts", controller.RenderPosts)
	}
	apiGroup := router.Group("/1.0")
	{
		apiGroup.GET("/me", controller.GetMe)
		apiGroup.GET("/me/sign_up", controller.SingUp)
		apiGroup.POST("/me/sign_in", controller.SignIn)

		apiGroup.GET("/posts", controller.Posts)

		apiGroup.GET("/user/:id", controller.User)
	}
	return router
}
