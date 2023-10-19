package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"
	"fmt"

	"github.com/tf63/code-api/api"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input api.NewTodo) (*api.Todo, error) {

	inputEntity := NewTodoDTO(input)

	todoId, err := r.Tr.CreateTodo(inputEntity)
	if err != nil {
		err = TodoErrorHandler("Create", err)
		return nil, err
	}

	todoEntity, err := r.Tr.ReadTodo(*todoId)
	if err != nil {
		err = TodoErrorHandler("Read", err)
		return nil, err
	}

	todo := TodoDTO(*todoEntity)

	return &todo, nil
}

// CreateProgramCode is the resolver for the createProgramCode field.
func (r *mutationResolver) CreateProgramCode(ctx context.Context, input api.NewProgramCode) (string, error) {
	panic(fmt.Errorf("not implemented: CreateProgramCode - createProgramCode"))
}

// Mutation returns api.MutationResolver implementation.
func (r *Resolver) Mutation() api.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
