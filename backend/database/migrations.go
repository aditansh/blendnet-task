package database

import (
	"github.com/aditansh/blendnet-task/be/models"
	"gorm.io/gorm"
)

func RunMigrations(DB *gorm.DB) {
	DB.AutoMigrate(&models.User{})
}
