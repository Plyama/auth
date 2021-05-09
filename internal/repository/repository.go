package repository

import (
	"github.com/plyama/auth/internal/models"
	"gorm.io/gorm"
)

type Users interface {
	Create(user models.User) error
}

type Repositories struct {
	Users Users
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Users: NewUsersRepository(db),
	}
}
