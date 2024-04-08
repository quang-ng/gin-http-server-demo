package repository

import (
	"todolist/infrastructure"
	"todolist/models"
)

// TodoRepository -> TodoRepository
type TodoRepository struct {
	db infrastructure.Database
}

// NewTodoRepository : fetching database
func NewTodoRepository(db infrastructure.Database) TodoRepository {
	return TodoRepository{
		db: db,
	}
}

// Save -> Method for saving todo to database
func (p TodoRepository) Save(todo models.Todo) error {
	return p.db.DB.Create(&todo).Error
}

// FindAll -> Method for fetching all todos from database
func (p TodoRepository) FindAll(todo models.Todo, keyword string) (*[]models.Todo, int64, error) {
	var todos []models.Todo
	var totalRows int64 = 0

	queryBuider := p.db.DB.Order("created_at desc").Model(&models.Todo{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			p.db.DB.Where("todo.title LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(todo).
		Find(&todos).
		Count(&totalRows).Error
	return &todos, totalRows, err
}

// Update -> Method for updating Post
func (p TodoRepository) Update(todo models.Todo) error {
	return p.db.DB.Save(&todo).Error
}

// Find -> Method for fetching post by id
func (p TodoRepository) Find(todo models.Todo) (models.Todo, error) {
	var todos models.Todo
	err := p.db.DB.
		Debug().
		Model(&models.Todo{}).
		Where(&todo).
		Take(&todos).Error
	return todos, err
}

// Delete Deletes Post
func (p TodoRepository) Delete(todo models.Todo) error {
	return p.db.DB.Delete(&todo).Error
}
