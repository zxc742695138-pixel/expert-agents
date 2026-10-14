package model

import "time"

type Task struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateTaskInput struct {
	Title string `json:"title" binding:"required,min=1,max=200"`
}

type UpdateTaskInput struct {
	Title *string `json:"title" binding:"omitempty,min=1,max=200"`
	Done  *bool   `json:"done"`
}
