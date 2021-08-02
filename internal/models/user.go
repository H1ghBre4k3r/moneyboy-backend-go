package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID            string `gorm:"primaryKey" json:"id"`
	Username      string `gorm:"unique" json:"username"`
	DisplayName   string `json:"displayName"`
	Password      string `json:"-"`
	Email         string `json:"email"`
	EmailVerified bool   `gorm:"type:boolean" json:"emailVerified"`
}
