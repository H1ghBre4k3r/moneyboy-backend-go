package database

import (
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/models"
	"gorm.io/gorm"
)

type UserConnection struct {
	db *gorm.DB
}

func createUserConnection(db *gorm.DB) *UserConnection {
	return &UserConnection{db}
}

// Find a user by username
func (uc *UserConnection) FindByUsername(username string) *models.User {
	user := &models.User{}
	if err := uc.db.Where("username = ?", username).First(user).Error; err != nil {
		return nil
	}
	return user
}

// Find a user by id
func (uc *UserConnection) FindById(id string) *models.User {
	user := &models.User{}
	if err := uc.db.Where("id = ?", id).First(user).Error; err != nil {
		return nil
	}
	return user
}

// Create a new user in the database
func (uc *UserConnection) Create(user *models.User) error {
	return uc.db.Create(user).Error
}

// Delete a user by id
func (uc *UserConnection) DeleteById(id string) error {
	return uc.db.Delete("id = ?", id).Error
}

// Delete a user by username
func (uc *UserConnection) DeleteByUsername(username string) error {
	return uc.db.Delete("username = ?", username).Error
}

// Update a user in the database
func (uc *UserConnection) Update(user *models.User) error {
	return uc.db.Save(user).Error
}
