package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/plyama/auth/internal/middlewares"
	"github.com/plyama/auth/internal/models"
	"github.com/plyama/auth/internal/requests"
	"github.com/plyama/auth/internal/responses"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// Create godoc
// @Summary Create a task
// @ID create-task
// @Accept  json
// @Param task body requests.NewTask true "NewTask info"
// @Param Authorization header string true "Insert your jwt"
// @Success 201 "NewTask created"
// @Failure 400,500
// @Router /tasks [post]
func (h *Task) Create(c *gin.Context) {
	req, err := requests.CreateTask(c.Request)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	user, err := middlewares.GetUserData(c.Request.Context())
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	if user.Role != models.Customer {
		c.Status(http.StatusForbidden)
		return
	}

	task := models.Task{
		CustomerID:  user.ID,
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

func (h *Task) GetTaskDetails(c *gin.Context) {
	req, err := requests.GetOne(c)
	if err != nil {
		log.Println(err)
		return
	}

	user, err := middlewares.GetUserData(c.Request.Context())
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	task, err := h.TaskService.GetDetails(req.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.Status(http.StatusBadRequest)
			return
		}
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	if task.CustomerID != user.ID && *task.CoachID != user.ID {
		c.Status(http.StatusForbidden)
		return
	}

	response := responses.GetTask(*task)
	c.JSON(http.StatusOK, response)
}

func (h *Task) GetTasks(c *gin.Context) {
	user, err := middlewares.GetUserData(c.Request.Context())
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	var taskModels *[]models.Task

	switch user.Role {
	case models.Customer:
		taskModels, err = h.TaskService.GetForCustomer(user.ID)
	case models.Coach:
		taskModels, err = h.TaskService.GetForCoach(user.ID)
	}
	if err != nil {
		log.Printf("failed to get tasks: %v", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	tasks := responses.GetTasks(*taskModels, responses.GetTaskPreview)
	c.JSON(http.StatusOK, *tasks)
}
