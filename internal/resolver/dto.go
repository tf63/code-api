package resolver

import (
	"strconv"

	"github.com/tf63/code-api/api/graph"
	"github.com/tf63/code-api/internal/entity"
)

func TodoDTO(e entity.Todo) graph.Todo {
	id := strconv.Itoa(int(e.Id))

	return graph.Todo{
		ID:    id,
		Title: e.Title,
		Done:  e.Done,
	}
}

func NewTodoDTO(s graph.NewTodo) entity.NewTodo {
	return entity.NewTodo{
		Title: s.Title,
	}
}
