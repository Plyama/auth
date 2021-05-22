package repository

import (
	"errors"
	"github.com/plyama/auth/internal/models"
	"gorm.io/gorm"
)

type TaskRepo struct {
	db *gorm.DB
}

func NewTasksRepo(db *gorm.DB) *TaskRepo {
	return &TaskRepo{
		db: db,
	}
}

func (r *TaskRepo) Create(task models.Task) error {
	tx := r.db.Create(&task)

	if task.ID == 0 {
		return ErrorDuplicate(errors.New("task with similar unique fields exists"))
	}

	return tx.Error
}

func (r *TaskRepo) GetAll() (*[]models.Task, error) {
	var tasks []models.Task
	tx := r.db.Find(&tasks)

	return &tasks, tx.Error
}
