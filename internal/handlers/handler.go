package handlers

import (
	"github.com/plyama/auth/internal/service"
)

type User struct {
	UserService service.Users
}

func NewUser(usersService service.Users) *User {
	return &User{
		UserService: usersService,
	}
}

type Handler struct {
	User *User
}
