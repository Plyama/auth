package service

import (
	"github.com/plyama/auth/internal/models"
	"github.com/plyama/auth/internal/repository"
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
