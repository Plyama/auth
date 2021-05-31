package internal

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/plyama/auth/internal/handlers"
	"github.com/plyama/auth/internal/middlewares"
	"github.com/plyama/auth/internal/services"
)

func NewRouter(services *services.Services) *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")

	r.Use(cors.New(config))

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
		customer.PUT("", middlewares.AuthorizeViaHeader, handler.User.Update)
		customer.GET("", middlewares.AuthorizeViaHeader, handler.User.GetMy)

		tasks := customer.Group("tasks")
		tasks.POST("", middlewares.AuthorizeViaHeader, handler.Task.Create)
		tasks.GET("/:id", middlewares.AuthorizeViaHeader, handler.Task.GetTaskDetails)
		tasks.GET("", middlewares.AuthorizeViaHeader, handler.Task.GetTasks)
	}

	coach := api.Group("coach")
	{
		coach.GET("google-callback", handler.User.SignUpWebCallback)
		coach.GET("", middlewares.AuthorizeViaHeader, handler.User.GetMy)
		coach.PUT("", middlewares.AuthorizeViaHeader, handler.User.Update)

		tasks := coach.Group("tasks")
		tasks.GET("/:id", middlewares.AuthorizeViaHeader, handler.Task.GetTaskDetails)
		tasks.GET("", middlewares.AuthorizeViaHeader, handler.Task.GetTasks)
	}

	return r
}
