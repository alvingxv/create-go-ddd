package dto

import "time"

type CreateTodoRequest struct {
	Title string `json:"title" valid:"required~Title cannot be empty"`
}
type CreateTodoResponse struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
