package mysql

import (
	"UAKI-WEB/entity"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(entity.User{})
}
