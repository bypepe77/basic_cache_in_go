package todo

import (
	"fmt"

	"github.com/bypepe77/basic_cache_in_go/pkg/cache"
	"github.com/bypepe77/basic_cache_in_go/pkg/models"
	"github.com/google/uuid"
)

type TodoController struct {
	todos []*models.Todo
	cache *cache.Cache
}

func NewTodoController() TodoInterface {
	return &TodoController{
		todos: []*models.Todo{},
		cache: cache.GetCache(),
	}
}

func (ctrl *TodoController) AddNewTodo() *models.Todo {
	todo := &models.Todo{
		ID:       uuid.New(),
		Content:  "todo content",
		Finished: false,
	}

	ctrl.todos = append(ctrl.todos, todo)
	// In case you want to set into cache the new created todo
	// ctrl.cache.SetTodo(todo)

	return todo
}

func (ctrl *TodoController) GetTodoById(todoID uuid.UUID) (*models.Todo, error) {
	todo, err := ctrl.cache.GetTodo(todoID)
	if err != nil {
		todo = getTodoFromList(ctrl.todos, todoID)
		if todo == nil {
			return nil, fmt.Errorf("Error retreiving todo")
		}
		ctrl.cache.SetTodo(todo)
	}
	return todo, nil
}

func getTodoFromList(todos []*models.Todo, todoID uuid.UUID) *models.Todo {
	for _, todo := range todos {
		if todo.ID == todoID {
			return todo
		}
	}
	return nil
}
