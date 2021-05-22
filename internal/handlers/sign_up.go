package handlers

import (
	"github.com/plyama/auth/internal/responses"
	"github.com/plyama/auth/internal/utils/auth"
	"log"
	"net/http"

	"github.com/plyama/auth/internal/models"
	"github.com/plyama/auth/internal/requests"
	"github.com/plyama/auth/internal/utils/oauth"

	"github.com/gin-gonic/gin"
)

// SignUpRedirect godoc
// @Summary Redirect to an OAuth page of Google
// @Description No data needed
// @ID sign-up-google
// @Success 307
// @Router /auth/google-oauth [get]
func (h *User) SignUpRedirect(c *gin.Context) {
	googleOauthConf := oauth.NewGoogleConfig(oauth.GoogleSignUpCallbackURL())
	redirectURL := googleOauthConf.AuthCodeURL("")

	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}

// SignUpCallback godoc
// @Summary Redirect to app's page with login
// @ID sign-up-google-callback
// @Success 200 {object} responses.JWT
// @Success 201 "User created"
// @Failure 400,500
// @Router /auth/google-oauth [get]
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
		resp := responses.JWT{
			Token:     jwt,
			TokenType: "jwt",
		}

		c.JSON(http.StatusOK, resp)
		return
	}

	user := models.User{
		Name:  userData.Name,
		Email: userData.Email,
		Role:  models.Customer,
	}

	err = h.UserService.Create(user)
	if err != nil {
		log.Printf("failed to create user in db: %v", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusCreated, "https://www.google.com/")
}
