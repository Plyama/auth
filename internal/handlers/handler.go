package handlers

import (
	"github.com/plyama/auth/internal/services"
)

type User struct {
	UserService services.UsersService
}

type Task struct {
	TaskService services.TasksService
}

func NewUser(usersService services.UsersService) *User {
	return &User{
		UserService: usersService,
	}
}

func NewTask(tasksService services.TasksService) *Task {
	return &Task{
		TaskService: tasksService,
	}
}

type Handler struct {
	User *User
	Task *Task
}
