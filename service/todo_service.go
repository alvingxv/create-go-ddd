package service

import (
	"github.com/alvingxv/todos-kelompok5/dto"
	"github.com/alvingxv/todos-kelompok5/entity"
	"github.com/alvingxv/todos-kelompok5/pkg/errs"
	"github.com/alvingxv/todos-kelompok5/repository/todo_repository"
	"github.com/asaskevich/govalidator"
)

type todoService struct {
	todoRepository todo_repository.TodoRepository
}

type TodoService interface {
	CreateTodo(payload dto.CreateTodoRequest) (*dto.CreateTodoResponse, errs.MessageErr)
}

func NewTodoService(todoRepository todo_repository.TodoRepository) TodoService {
	return &todoService{
		todoRepository: todoRepository,
	}
}

func (ts *todoService) CreateTodo(payload dto.CreateTodoRequest) (*dto.CreateTodoResponse, errs.MessageErr) {
	_, errv := govalidator.ValidateStruct(payload)

	if errv != nil {
		return nil, errs.NewBadRequest(errv.Error())
	}

	todos := &entity.Todo{
		Title: payload.Title,
	}

	err := ts.todoRepository.CreateTodo(todos)

	if err != nil {
		return nil, err
	}

	response := dto.CreateTodoResponse{
		Id:        todos.Id,
		Title:     todos.Title,
		Completed: todos.Completed,
		CreatedAt: todos.CreatedAt,
		UpdatedAt: todos.UpdatedAt,
	}

	return &response, nil

}
