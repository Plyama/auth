package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/plyama/auth/internal/middlewares"
	"github.com/plyama/auth/internal/models"
	"github.com/plyama/auth/internal/requests"
	"github.com/plyama/auth/internal/responses"
	"log"
	"net/http"
)

// Create godoc
// @Summary Create a task
// @ID create-task
// @Accept  json
// @Param task body requests.Task true "Task info"
// @Param Authorization header string true "Insert your jwt"
// @Success 201 "Task created"
// @Failure 400,500
// @Router /tasks [post]
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
		CustomerID:  customerID,
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

func (h *Task) GetAll(c *gin.Context) {
	taskModels, err := h.TaskService.GetAll()
	if err != nil {
		log.Printf("failed to get all tasks: %v", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	tasks := responses.GetTasks(*taskModels)
	c.JSON(http.StatusOK, *tasks)
}
