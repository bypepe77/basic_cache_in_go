package todo

import (
	"github.com/bypepe77/basic_cache_in_go/pkg/models"
	"github.com/google/uuid"
)

type TodoInterface interface {
	AddNewTodo() *models.Todo
	GetTodoById(todoID uuid.UUID) (*models.Todo, error)
}
