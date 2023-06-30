package todo_pg

import (
	"github.com/alvingxv/todos-kelompok5/entity"
	"github.com/alvingxv/todos-kelompok5/pkg/errs"
	"github.com/alvingxv/todos-kelompok5/repository/todo_repository"
	"gorm.io/gorm"
)

type todoPG struct {
	db *gorm.DB
}

func NewTodoPG(db *gorm.DB) todo_repository.TodoRepository {
	return &todoPG{
		db: db,
	}
}

func (t *todoPG) CreateTodo(todo *entity.Todo) errs.MessageErr {

	err := t.db.Create(&todo).Error

	if err != nil {
		return errs.NewInternalServerError("something Went Wrong")
	}

	return nil
}
