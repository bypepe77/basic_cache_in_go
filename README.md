<div align="center">
 <h1>Basic cache in Go</h1>
    <span><strong>Basic cache in go</strong> It's an example of how to use basic cache using Go</span><br />
</div>

## Installation
```bash
go get github.com/bypepe77/basic_cache_in_go

```
### How to run it?
```bash
 go run cmd/main.go

```

## Models
```go

 // Todo Model
type Todo struct {
	ID       uuid.UUID
	Content  string
	Finished bool
}

```

## How to use it

```go
todoCtrl := todo.NewTodoController()

todo := todoCtrl.AddNewTodo()

// Find todo(first call will not be cached)
todoNotCached, err := todoCtrl.GetTodoById(todo.ID)
if err != nil {
     fmt.Println("Error getting todo", err)
}
fmt.Println("todo not cached", todoNotCached)

// Find todo(cached)
todoCached, err := todoCtrl.GetTodoById(todo.ID)
if err != nil {
     fmt.Println("Error getting todo", err)
}
fmt.Println("todo cached", todoCached)
```




