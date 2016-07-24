package service

import (
	"github.com/gin-gonic/gin"

	postmodel "github.com/gotokatsuya/lboard/model/post"
	postparam "github.com/gotokatsuya/lboard/parameter/post"
)

func GetPosts(ctx *gin.Context, req *postparam.GetRequest) (*postparam.GetResponse, error) {
	postRepository := postmodel.NewRepository(ctx)
	postEntities, err := postRepository.Find()
	if err != nil {
		return nil, err
	}
	res := postparam.NewGetResponse()
	if err := res.Convert(postEntities); err != nil {
		return nil, err
	}
	return res, nil
}

func CreatePost() {

}
