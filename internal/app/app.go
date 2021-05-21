package app

import (
	"github.com/plyama/auth/internal/repository"
	"github.com/plyama/auth/internal/services"

	"gorm.io/gorm"
)

type App struct {
	Services *services.Services
}

func NewApp(db *gorm.DB) *App {
	return &App{
		Services: &services.Services{
			User: services.NewUsers(repository.NewUsersRepo(db)),
			Task: services.NewTasks(repository.NewTasksRepo(db)),
		},
	}
}
