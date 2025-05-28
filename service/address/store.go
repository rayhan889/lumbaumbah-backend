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
	return nil	
}

func (s *Store) GetAddressByUserID(id string) (types.Address, error) {
	return types.Address{}, nil
}
