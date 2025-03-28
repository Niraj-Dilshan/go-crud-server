package repository

import (
	"crud-server/data/request"
	"crud-server/model"
	"errors"
	"gorm.io/gorm"
)

type TaskRepositoryImpl struct {
	Db *gorm.DB
}

func NewTaskRepositoryImpl(Db *gorm.DB) TaskRepository {
	return &TaskRepositoryImpl{Db: Db}
}

func (t TaskRepositoryImpl) FindAll() (tasks []model.Task, err error) {
	results := t.Db.Find(&tasks)
	if results.Error != nil {
		return nil, results.Error
	}

	return tasks, nil
}

func (t TaskRepositoryImpl) FindById(taskId int) (task model.Task, err error) {
	result := t.Db.Find(&task, taskId)

	if result.Error != nil {
		return model.Task{}, result.Error
	}

	if result.RowsAffected == 0 {
		return model.Task{}, errors.New("Task is not found")
	}

	return task, nil
}

func (t *TaskRepositoryImpl) Save(task model.Task) error {
	result := t.Db.Create(&task)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (t *TaskRepositoryImpl) Update(task model.Task) error {
	var data = request.UpdateTaskRequest{
		Id:          task.Id,
		Name:        task.Name,
		Description: task.Description,
	}

	result := t.Db.Model(&task).Updates(data)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (t *TaskRepositoryImpl) Delete(taskId int) error {
	var task model.Task

	result := t.Db.Where("id = ?", taskId).Delete(&task)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
