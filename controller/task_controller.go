package controller

import (
	"crud-server/data/request"
	"crud-server/data/response"
	"crud-server/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TaskController struct {
	taskService service.TaskService
}

func NewTaskController(service service.TaskService) *TaskController {
	return &TaskController{taskService: service}
}

func (controller *TaskController) FindAll(ctx *gin.Context) {
	data, err := controller.taskService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   data,
	}

	ctx.JSON(http.StatusOK, res)
}

func (controller *TaskController) FindById(ctx *gin.Context) {
	taskId := ctx.Param("id")
	id, err := strconv.Atoi(taskId)

	data, err := controller.taskService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Code:    404,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   data,
	}
	ctx.JSON(http.StatusOK, res)
}

func (controller *TaskController) Create(ctx *gin.Context) {
	req := request.CreateTaskRequest{}
	ctx.ShouldBindJSON(&req)

	err := controller.taskService.Create(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, res)
}

func (controller *TaskController) Update(ctx *gin.Context) {
	req := request.UpdateTaskRequest{}
	err := ctx.ShouldBindJSON(&req)

	taskId := ctx.Param("id")
	id, err := strconv.Atoi(taskId)

	_, err = controller.taskService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Code:    404,
			Message: err.Error(),
		})
		return
	}

	req.Id = id

	err = controller.taskService.Update(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, res)
}

func (controller *TaskController) Delete(ctx *gin.Context) {
	taskId := ctx.Param("id")
	id, err := strconv.Atoi(taskId)

	_, err = controller.taskService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Code:    404,
			Message: err.Error(),
		})
		return
	}

	err = controller.taskService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, res)
}
