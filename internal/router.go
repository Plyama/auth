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

	api := r.Group("api/v1")

	authGroup := api.Group("auth")
	{
		authGroup.GET("google-oauth", handler.User.SignUpRedirect)
		authGroup.GET("google-callback", handler.User.SignUpCallback)
	}

	taskGroup := api.Group("tasks")
	{
		taskGroup.POST("", middlewares.Authorize, handler.Task.Create)
	}

	return r
}
