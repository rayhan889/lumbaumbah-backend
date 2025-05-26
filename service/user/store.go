package user

import (
	"database/sql"

	"github.com/rayhan889/lumbaumbah-backend/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateUser(user types.User) error {
	return nil
}