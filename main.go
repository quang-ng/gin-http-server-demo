package main

import (
	"todolist/api/controller"
	"todolist/api/repository"
	"todolist/api/routes"
	"todolist/api/service"
	"todolist/infrastructure"
	"todolist/models"
)

func init() {
	infrastructure.LoadEnv()
}

func main() {

	router := infrastructure.NewGinRouter()                     //router has been initialized and configured
	db := infrastructure.NewDatabase()                          // databse has been initialized and configured
	todoRepository := repository.NewTodoRepository(db)          // repository are being setup
	todoService := service.NewTodoService(todoRepository)       // service are being setup
	todoController := controller.NewTodoController(todoService) // controller are being set up
	todoRoute := routes.NewTodoRoute(todoController, router)    // todo routes are initialized
	todoRoute.Setup()                                           // todo routes are being setup

	db.DB.AutoMigrate(&models.Todo{}) // migrating todo model to datbase table
	router.Gin.Run(":8000")           //server started on 8000 port
}
