package models

import "time"

type Todo struct {
	ID          int64     `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `gorm:"size:200" json:"title"`
	Description string    `gorm:"size:3000" json:"description" `
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

// TableName method sets table name for Post model
func (todo *Todo) TableName() string {
	return "todo"
}

// ResponseMap -> response map method of Post
func (todo *Todo) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = todo.ID
	resp["title"] = todo.Title
	resp["description"] = todo.Description
	resp["created_at"] = todo.CreatedAt
	resp["updated_at"] = todo.UpdatedAt
	return resp
}
