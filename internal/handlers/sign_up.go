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
// @Router /google-oauth [get]
func (h *User) SignUpRedirect(c *gin.Context) {
	googleOauthConf := oauth.NewGoogleConfig(oauth.GoogleSignUpCallbackURL())
	redirectURL := googleOauthConf.AuthCodeURL("")

	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}

// SignUpWebCallback godoc
// @Summary Redirect to app's page with login
// @ID sign-up-google-callback
// @Success 200 {object} responses.JWT
// @Success 201 "User created"
// @Failure 400,500
// @Router /coach/google-oauth [get]
func (h *User) SignUpWebCallback(c *gin.Context) {
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

		c.SetCookie("token", jwt, 3600, "", "", false, false)
		c.JSON(http.StatusOK, resp)
		return
	}

	user := models.User{
		Name:  userData.Name,
		Email: userData.Email,
		Role:  models.Coach,
	}

	err = h.UserService.Create(user)
	if err != nil {
		log.Printf("failed to create user in db: %v", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusCreated, "https://www.google.com/")
}

// SignUpMobileCallback godoc
// @ID Sign-Up-Mobile-Callback
// @Param access_token header string true "Insert google token"
// @Success 201 "NewTask created"
// @Success 200 {object} responses.JWT
// @Failure 400,500
// @Router /customer/google-callback [post]
func (h *User) SignUpMobileCallback(c *gin.Context) {
	accessToken := c.GetHeader("access_token")
	if accessToken == "" {
		log.Printf("access token is: %s", accessToken)
		c.Status(http.StatusBadRequest)
		return
	}

	userData, err := oauth.GetGoogleUserInfo(accessToken)
	if err != nil {
		log.Printf("error while getting google user's info: %v\n", err)
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

	c.Status(http.StatusCreated)
}
