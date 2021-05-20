package handlers

import (
	"github.com/plyama/auth/internal/utils/auth"
	"log"
	"net/http"

	"github.com/plyama/auth/internal/models"
	"github.com/plyama/auth/internal/requests"
	"github.com/plyama/auth/internal/utils/oauth"

	"github.com/gin-gonic/gin"
)

func (h *User) SignUpRedirect(c *gin.Context) {
	googleOauthConf := oauth.NewGoogleConfig(oauth.GoogleSignUpCallbackURL())
	redirectURL := googleOauthConf.AuthCodeURL("")

	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}

func (h *User) SignUpCallback(c *gin.Context) {
	req, err := requests.CompleteOAuth(c.Request)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	userData, err := h.UserService.GetUserGoogleData(req.Code)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	registered, err := h.UserService.IsRegistered(userData.Email)
	if err != nil {
		log.Printf("failed to check if user is registered: %v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if registered {
		user, err := h.UserService.GetByEmail(userData.Email)
		if err != nil {
			log.Printf("failed to get user from db: %v", err)
			c.Status(http.StatusInternalServerError)
			return
		}

		jwt, err := auth.GenerateJWT(user)
		if err != nil {
			log.Printf("failed to generate jwt: %v", err)
			c.Status(http.StatusInternalServerError)
			return
		}

		c.String(http.StatusOK, jwt)
		return
	}

	user := models.User{
		Name:  userData.Name,
		Email: userData.Email,
	}

	err = h.UserService.UsersRepository.Create(user)
	if err != nil {
		log.Printf("failed to create user in db: %v", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusCreated, "https://www.google.com/")
}
