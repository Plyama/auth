package responses

import (
	"github.com/plyama/auth/internal/models"
)

type Task struct {
	ID          int               `json:"id"`
	CustomerID  int               `json:"customer_id,omitempty"`
	CoachID     *int              `json:"coach_id,omitempty"`
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	Status      models.TaskStatus `json:"status"`
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

func GetTaskPreview(model models.Task) Task {
	return Task{
		ID:     model.ID,
		Name:   model.Name,
		Status: model.Status,
	}
}

func GetTasks(models []models.Task, function func(task models.Task) Task) *[]Task {
	tasks := make([]Task, 0, len(models))
	for _, model := range models {
		tasks = append(tasks, function(model))
	}

	return &tasks
}
