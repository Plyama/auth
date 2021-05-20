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
		return ErrorDuplicate(errors.New("user with similar unique fields exists"))
	}

	return tx.Error
}

func (r *UserRepo) IsRegistered(email string) (bool, error) {
	var user models.User
	tx := r.db.Where("email = ?", email).First(&user)

	switch tx.Error {
	case gorm.ErrRecordNotFound:
		return false, nil
	case nil:
		return true, nil
	}

	return false, tx.Error
}

func (r *UserRepo) GetByEmail(email string) (models.User, error) {
	var user models.User
	tx := r.db.Where("email = ?", email).First(&user)

	return user, tx.Error
}
