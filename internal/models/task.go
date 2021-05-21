package models

type TaskStatus string

const (
	Published  = TaskStatus("published")
	InProgress = TaskStatus("in_progress")
	Completed  = TaskStatus("completed")
)

type Task struct {
	ID          int
	Customer    User
	CustomerID  int
	Coach       User
	CoachID     *int
	Name        string
	Description string
	Status      TaskStatus
}
