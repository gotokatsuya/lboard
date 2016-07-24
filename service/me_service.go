package service

import (
	"github.com/gin-gonic/gin"

	"github.com/gotokatsuya/lboard/library/net/context/me"
	meparam "github.com/gotokatsuya/lboard/parameter/me"
)

// GetMe return logined user
func GetMe(ctx *gin.Context, req *meparam.GetRequest) (*meparam.GetResponse, error) {
	res := meparam.NewGetResponse()
	if err := res.Convert(me.Get(ctx)); err != nil {
		return nil, err
	}
	return res, nil
}
