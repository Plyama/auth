package handlers

import (
	"net/http"

	"github.com/plyama/auth/internal/models"
	"github.com/plyama/auth/internal/requests"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) SignUp(c *gin.Context) {
	body, err := requests.NewSignUp(c.Request)
	if err != nil {
		// TODO: log error
		c.Status(http.StatusBadRequest)
		return
	}

	user := models.User{
		Name: body.Name,
		Email: body.Email,
		Password: body.Password,
	}
	err = h.UserService.CreateUser(user)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.Status(http.StatusCreated)
}
