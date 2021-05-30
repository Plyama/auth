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
	{
		api.GET("google-oauth", handler.User.SignUpRedirect)
	}

	customer := api.Group("customer")
	{
		customer.POST("google-callback", handler.User.SignUpMobileCallback)

		tasks := api.Group("tasks")
		tasks.POST("", middlewares.Authorize, handler.Task.Create)
		tasks.GET("", middlewares.Authorize, handler.Task.GetTasks)
	}

	coach := api.Group("coach")
	{
		coach.GET("google-callback", handler.User.SignUpWebCallback)
	}

	return r
}
