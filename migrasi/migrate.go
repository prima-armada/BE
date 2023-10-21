package migrasi

import (
	"par/domain/model"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Manager{})
	db.AutoMigrate(&model.Admin{})
	db.AutoMigrate(&model.Department{})
}
