package handler

import (
	"net/http"

	"github.com/alvingxv/create-go-ddd/dto"
	"github.com/alvingxv/create-go-ddd/pkg/errs"
	"github.com/alvingxv/create-go-ddd/service"
	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	todoService service.TodoService
}

func NewTodoHandler(todoService service.TodoService) todoHandler {
	return todoHandler{
		todoService: todoService,
	}
}
func (th *todoHandler) CreateTodo(ctx *gin.Context) {

	var todoRequest dto.CreateTodoRequest

	if err := ctx.ShouldBindJSON(&todoRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := th.todoService.CreateTodo(todoRequest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusCreated, result)

}
