package database

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/gotokatsuya/lboard/library/database"
)

// Get ...
func Get(ctx *gin.Context) *gorm.DB {
	value, exists := ctx.Get("database")
	if !exists || value == nil {
		db := database.NewGorm()
		if gin.IsDebugging() {
			db = db.Debug()
		}
		ctx.Set("database", db)
		return db
	}
	return value.(*gorm.DB)
}

// Begin ...
func Begin(ctx *gin.Context) *gorm.DB {
	txn := Get(ctx).Begin()
	ctx.Set("database", txn)
	return txn
}

// Commit ...
func Commit(ctx *gin.Context) {
	txn := Get(ctx).Commit()
	ctx.Set("database", txn)
}

// Rollback ...
func Rollback(ctx *gin.Context) {
	txn := Get(ctx).Rollback()
	ctx.Set("database", txn)
}
