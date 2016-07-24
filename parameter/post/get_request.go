package post

import (
	"github.com/gin-gonic/gin"
)

// GetRequest ...
type GetRequest struct {
}

// NewGetRequest ...
func NewGetRequest(ctx *gin.Context) (*GetRequest, error) {
	req := &GetRequest{}
	return req, nil
}
