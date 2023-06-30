package handler

import (
	"github.com/alvingxv/todos-kelompok5/database"
	"github.com/alvingxv/todos-kelompok5/repository/todo_repository/todo_pg"
	"github.com/alvingxv/todos-kelompok5/service"
	"github.com/gin-gonic/gin"
)

func StartApp() {

	database.HandleDatabaseConnection()
	db := database.GetDatabaseInstance()

	todoRepo := todo_pg.NewTodoPG(db)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := NewTodoHandler(todoService)

	// port := os.Getenv("PORT")
	port := "4000"
	r := gin.Default()

	todoRoute := r.Group("/todos")
	{
		todoRoute.POST("", todoHandler.CreateTodo)
	}
	r.Run("127.0.0.1:" + port)
}
