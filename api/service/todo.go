package service

import (
	"todolist/api/repository"
	"todolist/models"
)

// TodoService TodoService struct
type TodoService struct {
	repository repository.TodoRepository
}

// NewTodoService : returns the TodoService struct instance
func NewTodoService(r repository.TodoRepository) TodoService {
	return TodoService{
		repository: r,
	}
}

// Save -> calls todo repository save method
func (p TodoService) Save(todo models.Todo) error {
	return p.repository.Save(todo)
}

// FindAll -> calls todo repo find all method
func (p TodoService) FindAll(todo models.Todo, keyword string) (*[]models.Todo, int64, error) {
	return p.repository.FindAll(todo, keyword)
}

// Update -> calls todorepo update method
func (p TodoService) Update(todo models.Todo) error {
	return p.repository.Update(todo)
}

// Delete -> calls todo repo delete method
func (p TodoService) Delete(id int64) error {
	var todo models.Todo
	todo.ID = id
	return p.repository.Delete(todo)
}

// Find -> calls todo repo find method
func (p TodoService) Find(todo models.Todo) (models.Todo, error) {
	return p.repository.Find(todo)
}
