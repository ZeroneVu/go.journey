package business

import (
	"fmt"

	. "github.com/ZeroneVu/go.journey/minibar/models"
)

var currentId int

var ListTodos Todos

// Give us some seed data
func init() {
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
}

func RepoFindTodo(id int) Todo {
	for _, t := range ListTodos {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return Todo{}
}

//this is bad, I don't think it passes race condtions
func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	ListTodos = append(ListTodos, t)
	return t
}

func RepoDestroyTodo(id int) error {
	for i, t := range ListTodos {
		if t.Id == id {
			ListTodos = append(ListTodos[:i], ListTodos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
