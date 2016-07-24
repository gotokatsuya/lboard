package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index/main.tmpl", nil)
}
