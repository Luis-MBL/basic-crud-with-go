package migrations

import (
	"github.com/Luis-MBL/basic-crud-with-go/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
