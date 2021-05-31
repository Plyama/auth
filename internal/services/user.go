package services

import (
	"context"
	"github.com/pkg/errors"
	"github.com/plyama/auth/internal/models"
	"github.com/plyama/auth/internal/repository"
	"github.com/plyama/auth/internal/utils/oauth"
)

type Users struct {
	UsersRepository repository.Users
}

func NewUsers(usersRepo repository.Users) *Users {
	return &Users{
		UsersRepository: usersRepo,
	}
}

func (s *Users) Create(user models.User) error {
	return s.UsersRepository.Create(user)
}

func (s *Users) GetByID(ID int) (*models.User, error) {
	return s.UsersRepository.GetByID(ID)
}

func (s *Users) Update(user models.User) error {
	return s.UsersRepository.Update(user)
}

func (s *Users) IsRegistered(email string) (bool, error) {
	return s.UsersRepository.IsRegistered(email)
}

func (s *Users) GetUserGoogleData(oauthCode string) (*oauth.GoogleUserData, error) {
	conf := oauth.NewGoogleConfig(oauth.GoogleSignUpCallbackURL())
	token, err := conf.Exchange(context.Background(), oauthCode)
	if err != nil {
		return nil, errors.Wrap(err, "failed to exchange code for token")
	}

	userData, err := oauth.GetGoogleUserInfo(token.AccessToken)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user info using AccessToken")
	}

	return &userData, nil
}

func (s *Users) GetByEmail(email string) (models.User, error) {
	return s.UsersRepository.GetByEmail(email)
}
