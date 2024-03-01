package controllers

import (
	"fmt"
	"github.com/ducthang310/go-todo/src/config"
	"github.com/ducthang310/go-todo/src/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"net/http"
)

var db *gorm.DB = config.ConnectDB()

type todoRequest struct {
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
}

//type todoResponse struct {
//	todoRequest
//	ID uint `json:"id"`
//}

func CreateTodo(context *gin.Context) {
	var data todoRequest

	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{}
	todo.Description = data.Description
	todo.Complete = false

	result := db.Create(&todo)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	//var response todoResponse
	//response.ID = todo.ID
	//response.Description = todo.Description
	//response.Complete = todo.Complete

	context.JSON(http.StatusCreated, gin.H{
		"data": todo,
	})
}

func GetAllTodos(context *gin.Context) {
	var todos []models.Todo

	err := db.Find(&todos)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data": todos,
	})

}

func UpdateTodo(context *gin.Context) {
	var data todoRequest
	id := cast.ToUint(context.Param("id"))

	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{}

	todoById := db.Where("id = ?", id).First(&todo)
	if todoById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found"})
		return
	}

	todo.Description = data.Description
	todo.Complete = data.Complete

	result := db.Save(&todo)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	//var response todoResponse
	//response.ID = todo.ID
	//response.Description = todo.Description
	//response.Complete = todo.Complete

	context.JSON(http.StatusOK, gin.H{
		"data": todo,
		"error": nil,
	})
}

func GetTodoById(context *gin.Context) {
	id := cast.ToUint(context.Param("idTodo"))
	todo := models.Todo{}
	todoById := db.Where("id = ?", id).First(&todo)
	if todoById.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Todo does not exist"})
		return
	}
	context.JSON(http.StatusOK, todo)
}

func DeleteTodo(context *gin.Context) {
	todo := models.Todo{}
	id := cast.ToUint(context.Param("idTodo"))

	delete := db.Where("id = ?", id).Unscoped().Delete(&todo)
	fmt.Println(delete)

	context.JSON(http.StatusOK, nil)

}
