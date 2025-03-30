package dto

// CreateTaskRequest - Request DTO for creating a task
type CreateTaskRequest struct {
	Name        string `validate:"required,min=1,max=200" json:"name"`
	Description string `validate:"required,min=1,max=200" json:"description"`
}

// UpdateTaskRequest - Request DTO for updating a task
type UpdateTaskRequest struct {
	Id          int    `validate:"required"`
	Name        string `validate:"required,max=200,min=1" json:"name"`
	Description string `validate:"required,max=200,min=1" json:"description"`
}

// TaskResponse - Response DTO for task data
type TaskResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
