package service

import (
	"github.com/gin-gonic/gin"

	"github.com/gotokatsuya/lboard/library/net/context/me"
	postparam "github.com/gotokatsuya/lboard/parameter/post"
)

func GetPosts(ctx *gin.Context, req *postparam.GetRequest) (*postparam.GetResponse, error) {
	res := postparam.NewGetResponse()
	if err := res.Convert(me.Get(ctx)); err != nil {
		return nil, err
	}
	return res, nil
}

func CreatePost() {

}
