package services

import (
	"github.com/plyama/auth/internal/models"
	"github.com/plyama/auth/internal/utils/oauth"
)

type UsersService interface {
	Create(user models.User) error
	Update(user models.User) error
	GetByID(ID int) (*models.User, error)
	IsRegistered(email string) (bool, error)
	GetByEmail(email string) (models.User, error)
	GetUserGoogleData(oauthCode string) (*oauth.GoogleUserData, error)
}

type TasksService interface {
	Create(task models.Task) error
	GetDetails(taskID int) (*models.Task, error)
	GetForCustomer(userID int) (*[]models.Task, error)
	GetForCoach(userID int) (*[]models.Task, error)
}

type Services struct {
	User UsersService
	Task TasksService
}
