package service

import (
	"crud-server/data/request"
	"crud-server/data/response"
)

type TaskService interface {
	Create(task request.CreateTaskRequest) error
	Update(task request.UpdateTaskRequest) error
	Delete(taskId int) error
	FindById(taskId int) (task response.TaskResponse, err error)
	FindAll() (tasks []response.TaskResponse, err error)
}
