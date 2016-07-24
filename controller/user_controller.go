package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	userparam "github.com/gotokatsuya/lboard/parameter/user"
	"github.com/gotokatsuya/lboard/service"
)

func User(ctx *gin.Context)  {
	request, err := userparam.NewGetRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	res, err := service.GetUser(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"me": res.Instance})
	return
}
