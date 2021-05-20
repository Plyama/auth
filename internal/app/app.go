package app

import (
	"github.com/plyama/auth/internal/repository"
	"github.com/plyama/auth/internal/service"

	"gorm.io/gorm"
)

type App struct {
	Services *service.Services
}

func NewApp(db *gorm.DB) *App {
	return &App{
		Services: &service.Services{
			User: service.Users{
				UsersRepository: repository.NewUsersRepo(db),
			},
		},
	}
}
