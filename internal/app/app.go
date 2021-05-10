package app

import (
	"github.com/plyama/auth/internal/repository"

	"gorm.io/gorm"
)

type App struct {
	Repositories *repository.Repositories
}

func NewApp(db *gorm.DB) *App {
	return &App{
		Repositories: &repository.Repositories{
			Users: repository.NewUsersRepo(db),
		},
	}
}
