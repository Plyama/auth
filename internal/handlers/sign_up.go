package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/plyama/auth/internal/models"
	"github.com/plyama/auth/internal/repository"
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

	conf := oauth.NewGoogleConfig(oauth.GoogleSignUpCallbackURL())
	token, err := conf.Exchange(context.Background(), req.Code)
	if err != nil {
		log.Printf("failed to exchange code for token: %v", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	userData, err := oauth.GetGoogleUserInfo(token.AccessToken)
	if err != nil {
		log.Printf("failed to get user info using AccessToken: %v", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	user := models.User{
		Name:  userData.Name,
		Email: userData.Email,
	}

	err = h.UserService.UsersRepository.Create(user)
	if err != nil {
		log.Printf("failed to create user in db: %v", err)
		switch err.(type) {
		case repository.ErrorAlreadyExists:
			c.Status(http.StatusConflict)
		default:
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	c.Status(http.StatusCreated)
}
