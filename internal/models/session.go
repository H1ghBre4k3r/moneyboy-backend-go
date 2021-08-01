package models

import (
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	ID     string `gorm:"primaryKey"`
	UserID string `gorm:"size:255"`
	User   User
}
