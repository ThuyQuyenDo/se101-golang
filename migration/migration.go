package migration

import (
	usrentity "go-ecommerce/internal/user/business/entity"

	"gorm.io/gorm"
)

// only add column but cannot delete
func Migration(db *gorm.DB) {
	err := db.AutoMigrate(usrentity.User{})

	if err != nil {
		return
	}
}
