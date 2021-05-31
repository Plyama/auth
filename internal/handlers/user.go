package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/plyama/auth/internal/middlewares"
	"github.com/plyama/auth/internal/models"
	"github.com/plyama/auth/internal/requests"
	"github.com/plyama/auth/internal/responses"
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

func (h *User) GetMy(c *gin.Context) {
	user, err := middlewares.GetUserData(c.Request.Context())
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	userData, err := h.UserService.GetByID(user.ID)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	resp := responses.GetUser(*userData)

	c.JSON(http.StatusOK, resp)
}
