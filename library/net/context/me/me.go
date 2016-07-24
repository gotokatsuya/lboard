package me

import (
	"github.com/gin-gonic/gin"

	"github.com/gotokatsuya/lboard/model/user"
)

func Get(ctx *gin.Context) *user.Entity {
	value, exists := ctx.Get("me")
	if !exists {
		return nil
	}
	return value.(*user.Entity)
}

func Set(ctx *gin.Context, me *user.Entity) {
	ctx.Set("me", me)
}
