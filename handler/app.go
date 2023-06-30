package handler

import (
	"github.com/alvingxv/create-go-ddd/database"
	"github.com/alvingxv/create-go-ddd/repository/todo_repository/todo_pg"
	"github.com/alvingxv/create-go-ddd/service"
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
