package main

import (
	"fmt"

	"github.com/bypepe77/basic_cache_in_go/pkg/todo"
)

func main() {
	todoCtrl := todo.NewTodoController()

	todo := todoCtrl.AddNewTodo()

	// Find todo(first call will not be cached)
	todoNotCached, err := todoCtrl.GetTodoById(todo.ID)
	if err != nil {
		fmt.Println("Error getting todo", err)
	}
	fmt.Println("todo not cached", todoNotCached)
	// Response:
	//

	// Find todo(cached)
	todoCached, err := todoCtrl.GetTodoById(todo.ID)
	if err != nil {
		fmt.Println("Error getting todo", err)
	}
	fmt.Println("todo cached", todoCached)

}
