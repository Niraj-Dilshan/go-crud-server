package task

import (
	"crud-server/api/task/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Response defines the standard API response format
type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

// ErrorResponse defines the standard API error response format
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Controller handles HTTP requests for tasks
type Controller struct {
	service *Service
}

// NewController creates a new task controller
func NewController(service *Service) *Controller {
	return &Controller{service: service}
}

// RegisterRoutes registers all task-related routes to the provided router group
func (c *Controller) RegisterRoutes(rg *gin.RouterGroup) {
	tasks := rg.Group("/tasks")
	tasks.GET("", c.FindAll)
	tasks.GET("/:id", c.FindById)
	tasks.POST("", c.Create)
	tasks.PATCH("/:id", c.Update)
	tasks.DELETE("/:id", c.Delete)
}

// FindAll handles GET /tasks
func (c *Controller) FindAll(ctx *gin.Context) {
	data, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   data,
	})
}

// FindById handles GET /tasks/:id
func (c *Controller) FindById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID format",
		})
		return
	}

	task, err := c.service.FindById(id)
	if err != nil {
		if err == ErrTaskNotFound {
			ctx.JSON(http.StatusNotFound, ErrorResponse{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   task,
	})
}

// Create handles POST /tasks
func (c *Controller) Create(ctx *gin.Context) {
	var req dto.CreateTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if err := c.service.Create(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, Response{
		Code:   http.StatusCreated,
		Status: "success",
	})
}

// Update handles PATCH /tasks/:id
func (c *Controller) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID format",
		})
		return
	}

	var req dto.UpdateTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	req.Id = id

	if err := c.service.Update(req); err != nil {
		if err == ErrTaskNotFound {
			ctx.JSON(http.StatusNotFound, ErrorResponse{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, Response{
		Code:   http.StatusOK,
		Status: "success",
	})
}

// Delete handles DELETE /tasks/:id
func (c *Controller) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID format",
		})
		return
	}

	if err := c.service.Delete(id); err != nil {
		if err == ErrTaskNotFound {
			ctx.JSON(http.StatusNotFound, ErrorResponse{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, Response{
		Code:   http.StatusOK,
		Status: "success",
	})
}
