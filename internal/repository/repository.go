package repository

import (
	"github.com/plyama/auth/internal/models"

	"gorm.io/gorm"
)

type Users interface {
	Create(user models.User) error
	IsRegistered(email string) (bool, error)
	GetByEmail(email string) (models.User, error)
}

type Tasks interface {
	Create(task models.Task) error
}

type Repositories struct {
	Users Users
	Tasks Tasks
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Users: NewUsersRepo(db),
		Tasks: NewTasksRepo(db),
	}
}
