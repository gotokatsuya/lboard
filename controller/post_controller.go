package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	postparam "github.com/gotokatsuya/lboard/parameter/post"
	"github.com/gotokatsuya/lboard/service"
)

func Posts(ctx *gin.Context) {
	request, err := postparam.NewGetRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	res, err := service.GetPosts(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"posts": res.Instances})
	return
}

func RenderPosts(ctx *gin.Context) {
	request, err := postparam.NewGetRequest(ctx)
	if err != nil {
		return
	}
	res, err := service.GetPosts(ctx, request)
	if err != nil {
		return
	}
	ctx.HTML(http.StatusOK, "posts/list.tmpl", gin.H{"Posts": res.Instances})
	return
}
