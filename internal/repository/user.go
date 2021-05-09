package repository

import (
	"github.com/plyama/auth/internal/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Create(user models.User) error {
	tx := r.db.Create(user)
	return tx.Error
}
