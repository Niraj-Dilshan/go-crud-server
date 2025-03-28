package router

import (
	"crud-server/controller"
	"github.com/gin-gonic/gin"
)

func TaskRouter(taskController *controller.TaskController) *gin.Engine {
	service := gin.Default()

	router := service.Group("/tasks")

	router.GET("", taskController.FindAll)
	router.GET("/:id", taskController.FindById)
	router.POST("", taskController.Create)
	router.PATCH("/:id", taskController.Update)
	router.DELETE("/:id", taskController.Delete)

	return service
}
