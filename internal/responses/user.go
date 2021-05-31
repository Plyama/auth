package responses

import (
	"github.com/plyama/auth/internal/models"
)

type User struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	PicURL *string `json:"picture_url"`
}

func GetUser(model models.User) User {
	return User{
		ID:     model.ID,
		Name:   model.Name,
		Email:  model.Email,
		PicURL: model.PicURL,
	}
}
