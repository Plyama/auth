package internal

import (
	"github.com/plyama/auth/internal/handlers"
	"github.com/plyama/auth/internal/service"

	"github.com/gin-gonic/gin"
)

func NewRouter(services *service.Services) *gin.Engine {
	r := gin.Default()

	handler := handlers.Handler{
		User: handlers.NewUser(services.User),
	}

	r.GET("google-oauth", handler.User.SignUpRedirect)
	r.GET("google-callback", handler.User.SignUpCallback)

	return r
}
