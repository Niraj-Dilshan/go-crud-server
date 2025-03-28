package main

import (
	"crud-server/config"
	"crud-server/controller"
	"crud-server/repository"
	"crud-server/router"
	"crud-server/service"
	"crud-server/utils"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"time"
)

func main() {
	//Database
	db := config.DatabaseConnection()
	validate := validator.New()

	//Init Repository
	taskRepository := repository.NewTaskRepositoryImpl(db)

	//Init Service
	taskService, err := service.NewTaskServiceImpl(taskRepository, validate)
	if err != nil {
		// Handle error appropriately, such as logging and exiting
		log.Fatalf("Error initializing task service: %v", err)
	}

	//Init controller
	taskController := controller.NewTaskController(taskService)

	//Router
	routes := router.TaskRouter(taskController)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err = server.ListenAndServe()
	utils.ErrorPanic(err)

}
