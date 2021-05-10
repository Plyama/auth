package service

import (
	"github.com/plyama/auth/internal/models"
	"github.com/plyama/auth/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	UsersRepository repository.Users
}

func NewUsers(usersRepo repository.Users) *Users {
	return &Users{
		UsersRepository: usersRepo,
	}
}

func (s *Users) CreateUser(user models.User) error {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		// TODO: log error
		return err
	}

	user.Password = string(passwordBytes)

	return s.UsersRepository.Create(user)
}
