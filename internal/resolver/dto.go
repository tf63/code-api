package resolver

import (
	"strconv"

	"github.com/tf63/code-api/api"
	"github.com/tf63/code-api/internal/entity"
)

func TodoDTO(e entity.Todo) api.Todo {
	id := strconv.Itoa(int(e.Id))

	return api.Todo{
		ID:    id,
		Title: e.Title,
		Done:  e.Done,
	}
}

func NewTodoDTO(s api.NewTodo) entity.NewTodo {
	return entity.NewTodo{
		Title: s.Title,
	}
}
