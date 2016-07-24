package service

import (
	"github.com/gin-gonic/gin"

	"github.com/gotokatsuya/lboard/library/net/context/me"
	userparam "github.com/gotokatsuya/lboard/parameter/user"
)

// GetMe return logined user
func GetUser(ctx *gin.Context, req *userparam.GetRequest) (*userparam.GetResponse, error) {
	res := userparam.NewGetResponse()
	if err := res.Convert(me.Get(ctx)); err != nil {
		return nil, err
	}
	return res, nil
}
