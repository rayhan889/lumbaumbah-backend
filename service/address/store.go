package address

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

func (s *Store) CreateAddress(address types.Address) error {
	result := s.db.Create(&address)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Store) GetAddressesByUserID(id string) ([]types.Address, error) {
	var addresses []types.Address

	result := s.db.Where("user_id = ?", id).Find(&addresses)
	if result.Error != nil {
		return addresses, result.Error
	}

	return addresses, nil
}
