package database

import (
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/models"
	"gorm.io/gorm"
)

type SessionConnection struct {
	db *gorm.DB
}

func createSessionConnection(db *gorm.DB) *SessionConnection {
	return &SessionConnection{db}
}

func (sc *SessionConnection) Create(session *models.Session) error {
	return sc.db.Create(session).Error
}

func (sc *SessionConnection) Get(id string) *models.Session {
	session := &models.Session{}
	if err := sc.db.Where("id = ?", id).First(session).Error; err != nil {
		return nil
	}
	return session
}

func (sc *SessionConnection) Delete(id string) error {
	return sc.db.Where("id = ?", id).Delete(&models.Session{}).Error
}
