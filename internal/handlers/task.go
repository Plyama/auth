package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/plyama/auth/internal/middlewares"
	"github.com/plyama/auth/internal/models"
	"github.com/plyama/auth/internal/requests"
	"log"
	"net/http"
)

func (h *Task) Create(c *gin.Context) {
	req, err := requests.NewTask(c.Request)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	customerID, err := middlewares.GetUserID(c.Request.Context())
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	task := models.Task{
		CustomerID:  int(customerID),
		Name:        req.Name,
		Description: req.Description,
	}

	err = h.TaskService.Create(task)
	if err != nil {
		log.Println("failed to create task")
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}
