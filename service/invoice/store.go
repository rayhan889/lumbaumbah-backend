package invoice

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

func (s *Store) GetInvoiceByID(id string) (types.Invoice, error) {
	var invoice types.Invoice
	result := s.db.Where("id = ?", id).First(&invoice)
	if result.Error != nil {
		return invoice, result.Error
	}

	return invoice, nil
}
