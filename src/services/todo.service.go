package services

import (
	"errors"
	"fmt"
	"github.com/ducthang310/go-todo/src/models"
	"gorm.io/gorm"
)

type CreateTodoDTO struct {
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
}

type UpdateTodoDTO struct {
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
}

type TodoService struct {
	db       *gorm.DB
	Complete bool
}

func (s *TodoService) Create(dto CreateTodoDTO) (models.Todo, error) {
	todo := models.Todo{}
	todo.Description = dto.Description
	todo.Complete = false

	result := s.db.Create(&todo)
	if result.Error != nil {
		return models.Todo{}, errors.New("something went wrong")
	}
	return todo, nil
}

func (s *TodoService) Update(id int, dto UpdateTodoDTO) (models.Todo, error) {
	todo := models.Todo{}

	todoById := s.db.Where("id = ?", id).First(&todo)
	if todoById.Error != nil {
		return models.Todo{}, errors.New("todo does not exist")
	}

	todo.Description = dto.Description
	todo.Complete = dto.Complete

	result := s.db.Save(&todo)
	if result.Error != nil {
		return models.Todo{}, errors.New("something went wrong")
	}
	return todo, nil
}

func (s *TodoService) GetTodoById(id int) (models.Todo, error) {
	todo := models.Todo{}
	todoById := s.db.Where("id = ?", id).First(&todo)
	if todoById.Error != nil {
		return models.Todo{}, errors.New("todo does not exist")
	}

	return todo, nil
}

func (s *TodoService) GetAllTodos() ([]models.Todo, error) {
	var todos []models.Todo

	err := s.db.Find(&todos)
	if err.Error != nil {
		return nil, errors.New("something went wrong")
	}

	return todos, nil
}

func (s *TodoService) Delete(id int) (bool, error) {
	todo := models.Todo{}
	delete := s.db.Where("id = ?", id).Unscoped().Delete(&todo)
	fmt.Println(delete)

	return true, nil
}
