package task

import (
	"crud-server/api/task/dto"
	"crud-server/api/task/model"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Service handles task business logic
type Service struct {
	db       *gorm.DB
	validate *validator.Validate
}

// NewService creates a new task service
func NewService(db *gorm.DB, validate *validator.Validate) *Service {
	return &Service{
		db:       db,
		validate: validate,
	}
}

// FindAll retrieves all tasks
func (s *Service) FindAll() ([]dto.TaskResponse, error) {
	var tasks []model.Task
	result := s.db.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}

	var responses []dto.TaskResponse
	for _, task := range tasks {
		responses = append(responses, dto.TaskResponse{
			Id:          task.Id,
			Name:        task.Name,
			Description: task.Description,
		})
	}

	return responses, nil
}

// FindById retrieves a task by ID
func (s *Service) FindById(taskId int) (dto.TaskResponse, error) {
	var task model.Task
	result := s.db.Find(&task, taskId)

	if result.Error != nil {
		return dto.TaskResponse{}, result.Error
	}

	if result.RowsAffected == 0 {
		return dto.TaskResponse{}, ErrTaskNotFound
	}

	return dto.TaskResponse{
		Id:          task.Id,
		Name:        task.Name,
		Description: task.Description,
	}, nil
}

// Create creates a new task
func (s *Service) Create(req dto.CreateTaskRequest) error {
	if err := s.validate.Struct(req); err != nil {
		return err
	}

	task := model.Task{
		Name:        req.Name,
		Description: req.Description,
	}

	result := s.db.Create(&task)
	return result.Error
}

// Update updates an existing task
func (s *Service) Update(req dto.UpdateTaskRequest) error {
	if err := s.validate.Struct(req); err != nil {
		return err
	}

	result := s.db.Model(&model.Task{}).Where("id = ?", req.Id).Updates(map[string]interface{}{
		"name":        req.Name,
		"description": req.Description,
	})

	if result.RowsAffected == 0 {
		return ErrTaskNotFound
	}

	return result.Error
}

// Delete removes a task by ID
func (s *Service) Delete(taskId int) error {
	result := s.db.Delete(&model.Task{}, taskId)

	if result.RowsAffected == 0 {
		return ErrTaskNotFound
	}

	return result.Error
}
