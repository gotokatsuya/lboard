package route

import (
	"github.com/gin-gonic/gin"

	"github.com/gotokatsuya/lboard/controller"
)

// Init ...
func Init() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("view/**/*")
	renderGroup := router.Group("/render")
	{
		renderGroup.GET("/posts", controller.RenderPosts)
	}
	apiGroup := router.Group("/1.0")
	{
		apiGroup.GET("/me", controller.GetMe)
		apiGroup.GET("/me/sign_up", controller.Sing_up)
		apiGroup.POST("/me/login", controller.Login)

		apiGroup.GET("/posts", controller.Posts)

		apiGroup.GET("/user/:id", controller.User)
	}
	return router
}
