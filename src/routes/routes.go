package routes

import (
	"github.com/ducthang310/go-todo/src/controllers"
	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()

	router.GET("/todos", controllers.GetAllTodos)
	router.POST("/todos", controllers.CreateTodo)
	router.GET("/todos/:id", controllers.GetTodoById)
	router.PATCH("/todos/:id", controllers.UpdateTodo)
	router.DELETE("/todos/:id", controllers.DeleteTodo)

	router.Run()
}
