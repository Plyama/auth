package services

import (
	"github.com/plyama/auth/internal/models"
	"github.com/plyama/auth/internal/repository"
)

type Tasks struct {
	TasksRepository repository.Tasks
}

func NewTasks(tasksRepo repository.Tasks) *Tasks {
	return &Tasks{
		TasksRepository: tasksRepo,
	}
}

func (s *Tasks) Create(task models.Task) error {
	task.Status = models.Published
	return s.TasksRepository.Create(task)
}
