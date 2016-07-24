package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	meparam "github.com/gotokatsuya/lboard/parameter/me"
	"github.com/gotokatsuya/lboard/service"
)

// 自分のデータ取得
func GetMe(ctx *gin.Context) {
	request, err := meparam.NewGetRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	res, err := service.GetMe(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"me": res.Instance})
	return
}

func SingUp(ctx *gin.Context) {

}

func SignIn(ctx *gin.Context) {

}
