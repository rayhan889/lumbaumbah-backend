package admin

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

func (s *Store) CreateAdmin(admin types.Admin) error {
	result := s.db.Create(&admin)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Store) GetAdminByID(id string) (types.Admin, error) {
	var admin types.Admin
	result := s.db.Where("id = ?", id).First(&admin)
	if result.Error != nil {
		return admin, result.Error
	}

	return admin, nil
}

func (s *Store) GetAdminByEmail(email string) (types.Admin, error) {
	var admin types.Admin
	result := s.db.Where("email = ?", email).First(&admin)
	if result.Error != nil {
		return admin, result.Error
	}

	return admin, nil
}