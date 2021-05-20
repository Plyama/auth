package repository

import (
	"errors"

	"github.com/plyama/auth/internal/models"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Create(user models.User) error {
	tx := r.db.Create(&user)

	if user.ID == 0 {
		return ErrorAlreadyExists(errors.New("user with similar unique fields exists"))
	}

	return tx.Error
}
