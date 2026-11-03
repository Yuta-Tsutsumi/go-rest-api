package repository

import "go-rest-api/model"

type ITaskRepository interface {
	GetAllTasks(tasks *[]model.Task, userId uint) error
}