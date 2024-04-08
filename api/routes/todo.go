package routes

import (
	"todolist/api/controller"
	"todolist/infrastructure"
)

// PostRoute -> Route for question module
type TodoRoute struct {
	Controller controller.TodoController
	Handler    infrastructure.GinRouter
}

// NewTodoRoute -> initializes new choice rouets
func NewTodoRoute(
	controller controller.TodoController,
	handler infrastructure.GinRouter,

) TodoRoute {
	return TodoRoute{
		Controller: controller,
		Handler:    handler,
	}
}

// Setup -> setups new choice Routes
func (p TodoRoute) Setup() {
	todo := p.Handler.Gin.Group("/todos") //Router group
	{
		todo.GET("/", p.Controller.GetTodos)
		todo.POST("/", p.Controller.AddTodo)
		todo.GET("/:id", p.Controller.GetTodo)
		todo.DELETE("/:id", p.Controller.DeleteTodo)
		todo.PUT("/:id", p.Controller.UpdateTodo)
	}
}
