package route

import (
	"github.com/gin-gonic/gin"

	"github.com/gotokatsuya/lboard/controller"
)

// Init ...
func Init() *gin.Engine {
	router := gin.Default()
	apiGroup := router.Group("/1.0")
	{
		apiGroup.GET("/me", controller.GetMe)
		apiGroup.GET("/me/sign_up", controller.Sing_up)
		apiGroup.POST("/me/login", controller.Login)

		apiGroup.GET("/posts", controller.Posts)
		apiGroup.POST("/post", controller.Post)

		apiGroup.GET("/user/:id", controller.User)
	}
	return router
}
