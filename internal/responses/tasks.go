package responses

import (
	"github.com/plyama/auth/internal/models"
)

type Task struct {
	ID          int
	CustomerID  int
	CoachID     *int
	Name        string
	Description string
	Status      models.TaskStatus
}

func GetTask(model models.Task) Task {
	return Task{
		ID:          model.ID,
		CustomerID:  model.CustomerID,
		CoachID:     model.CoachID,
		Name:        model.Name,
		Description: model.Description,
		Status:      model.Status,
	}
}

func GetTasks(models []models.Task) *[]Task {
	tasks := make([]Task, 0, len(models))
	for _, model := range models {
		tasks = append(tasks, GetTask(model))
	}

	return &tasks
}
