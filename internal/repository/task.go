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

func (r *TaskRepo) GetDetails(taskID int) (*models.Task, error) {
	var task models.Task
	tx := r.db.Where("id = ?", taskID).First(&task)

	return &task, tx.Error
}

func (r *TaskRepo) GetForCustomer(userID int) (*[]models.Task, error) {
	var tasks []models.Task
	tx := r.db.Where("customer_id = ?", userID).Find(&tasks)

	return &tasks, tx.Error
}

func (r *TaskRepo) GetForCoach(userID int) (*[]models.Task, error) {
	var tasks []models.Task
	tx := r.db.Where("coach_id = ?", userID).Find(&tasks)

	return &tasks, tx.Error
}
