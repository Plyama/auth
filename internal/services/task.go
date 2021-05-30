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

func (s *Tasks) GetDetails(taskID int) (*models.Task, error) {
	return s.TasksRepository.GetDetails(taskID)
}

func (s *Tasks) GetForCustomer(userID int) (*[]models.Task, error) {
	return s.TasksRepository.GetForCustomer(userID)
}

func (s *Tasks) GetForCoach(userID int) (*[]models.Task, error) {
	return s.TasksRepository.GetForCoach(userID)
}
