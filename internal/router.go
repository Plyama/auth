package internal

import (
	"github.com/plyama/auth/internal/handlers"
	"github.com/plyama/auth/internal/middlewares"
	"github.com/plyama/auth/internal/services"

	"github.com/gin-gonic/gin"
)

func NewRouter(services *services.Services) *gin.Engine {
	r := gin.Default()

	handler := handlers.Handler{
		User: handlers.NewUser(services.User),
		Task: handlers.NewTask(services.Task),
	}

	r.GET("google-oauth", handler.User.SignUpRedirect)
	r.GET("google-callback", handler.User.SignUpCallback)
	r.POST("create-task", middlewares.Authorize, handler.Task.Create)

	return r
}
