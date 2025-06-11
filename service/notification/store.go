package notification

import (
	"github.com/rayhan889/lumbaumbah-backend/types"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateNotification(notification types.Notification) error {
	result := s.db.Create(&notification)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
