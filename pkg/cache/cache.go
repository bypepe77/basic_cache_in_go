package cache

import (
	"fmt"
	"time"

	"github.com/bypepe77/basic_cache_in_go/pkg/models"
	"github.com/google/uuid"
)

type Cache struct {
	todo map[uuid.UUID]*CachedTodo
	ttl  time.Duration
}

type CachedTodo struct {
	todo      *models.Todo
	createdAt time.Time
}

var instance *Cache

func GetCache() *Cache {
	if instance == nil {
		instance = &Cache{
			todo: make(map[uuid.UUID]*CachedTodo),
			ttl:  time.Hour,
		}
	}
	return instance
}

func (c *Cache) SetTodo(todoModel *models.Todo) {
	cachedTodo := &CachedTodo{
		todo:      todoModel,
		createdAt: time.Now(),
	}
	c.todo[cachedTodo.todo.ID] = cachedTodo
}

func (c *Cache) GetTodo(todoID uuid.UUID) (*models.Todo, error) {
	cachedTodo, ok := c.todo[todoID]
	if !ok {
		fmt.Println("Not cached...")
		return nil, fmt.Errorf("Todo not found in cache")
	}
	if time.Since(cachedTodo.createdAt) > c.ttl {
		delete(c.todo, todoID)

		return nil, fmt.Errorf("Todo expired from cache")
	}

	fmt.Println("Cached response!")
	return cachedTodo.todo, nil
}

func (c *Cache) DeleteTodo(todoID uuid.UUID) {
	delete(c.todo, todoID)
}
