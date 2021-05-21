package services

import (
	"github.com/plyama/auth/internal/models"
	"github.com/plyama/auth/internal/utils/oauth"
)

type UsersService interface {
	Create(user models.User) error
	IsRegistered(email string) (bool, error)
	GetByEmail(email string) (models.User, error)
	GetUserGoogleData(oauthCode string) (*oauth.GoogleUserData, error)
}

type TasksService interface {
	Create(task models.Task) error
}

type Services struct {
	User UsersService
	Task TasksService
}
