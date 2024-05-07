package database

import (
	"github.com/Daanooo/vams/internal/model"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&model.AirlineInfo{})
}
