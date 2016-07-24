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
	}
	return router
}
