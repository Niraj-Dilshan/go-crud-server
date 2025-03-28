package service

import (
	"crud-server/data/request"
	"crud-server/data/response"
	"crud-server/model"
	"crud-server/repository"
	"errors"
	"github.com/go-playground/validator/v10"
)

type TaskServiceImpl struct {
	TaskRepository repository.TaskRepository
	Validate       *validator.Validate
}

func NewTaskServiceImpl(taskRepository repository.TaskRepository, validate *validator.Validate) (service TaskService, err error) {
	if validate == nil {
		return nil, errors.New("validator instance cannot be nil")
	}
	return &TaskServiceImpl{
		TaskRepository: taskRepository,
		Validate:       validate,
	}, err
}

func (t TaskServiceImpl) FindAll() (tasks []response.TaskResponse, err error) {
	result, err := t.TaskRepository.FindAll()
	if err != nil {
		return nil, err
	}

	for _, value := range result {
		task := response.TaskResponse{
			Id:          value.Id,
			Name:        value.Name,
			Description: value.Description,
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (t TaskServiceImpl) FindById(taskId int) (response.TaskResponse, error) {
	data, err := t.TaskRepository.FindById(taskId)
	if err != nil {
		return response.TaskResponse{}, err
	}

	res := response.TaskResponse{
		Id:          data.Id,
		Name:        data.Name,
		Description: data.Description,
	}
	return res, nil
}

func (t *TaskServiceImpl) Create(task request.CreateTaskRequest) (err error) {
	err = t.Validate.Struct(task)

	if err != nil {
		return err
	}

	m := model.Task{
		Name:        task.Name,
		Description: task.Description,
	}

	t.TaskRepository.Save(m)

	return nil
}

func (t *TaskServiceImpl) Update(task request.UpdateTaskRequest) (err error) {
	data, err := t.TaskRepository.FindById(task.Id)

	if err != nil {
		return err
	}

	data.Name = task.Name
	data.Description = task.Description
	t.TaskRepository.Update(data)
	return nil
}

func (t *TaskServiceImpl) Delete(taskId int) (err error) {
	err = t.TaskRepository.Delete(taskId)

	if err != nil {
		return err
	}

	return nil
}
