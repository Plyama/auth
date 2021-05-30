package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/plyama/auth/internal/middlewares"
	"github.com/plyama/auth/internal/models"
	"github.com/plyama/auth/internal/requests"
	"log"
	"net/http"
)

func (h *User) Update(c *gin.Context) {
	req, err := requests.NewUpdateUser(c.Request)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	user, err := middlewares.GetUserData(c.Request.Context())
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	userUpdate := models.User{
		ID:     user.ID,
		PicURL: &req.PicURL,
		DOB:    &req.DOB,
	}

	err = h.UserService.Update(userUpdate)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
