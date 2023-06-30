package todo_repository

import (
	"github.com/alvingxv/create-go-ddd/entity"
	"github.com/alvingxv/create-go-ddd/pkg/errs"
)

type TodoRepository interface {
	CreateTodo(todo *entity.Todo) errs.MessageErr
}
