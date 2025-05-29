package laundry

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

func (s *Store) CreateLaundryType(laundryType types.LaundryType) error {
	result := s.db.Create(&laundryType)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Store) GetLaundryTypes() ([]types.LaundryType, error) {
	var types []types.LaundryType
	result := s.db.Find(&types)
	if result.Error != nil {
		return types, result.Error
	}

	return types, nil
}