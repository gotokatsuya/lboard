package scheme

import (
	"github.com/jinzhu/gorm"

	postmodel "github.com/gotokatsuya/lboard/model/post"
	usermodel "github.com/gotokatsuya/lboard/model/user"
)

func createTables(db *gorm.DB, models []interface{}) {
	for _, model := range models {
		db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4").CreateTable(model)
	}
}

func autoMigrateTables(db *gorm.DB, models []interface{}) {
	for _, model := range models {
		db.AutoMigrate(model)
	}
}

func Sync(db *gorm.DB) {
	models := []interface{}{
		new(usermodel.Entity),
		new(postmodel.Entity),
	}
	createTables(db, models)
	autoMigrateTables(db, models)
}
