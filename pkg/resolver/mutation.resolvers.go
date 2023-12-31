package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"
	"fmt"

	"github.com/tf63/code-api/api"
)

// CreateFrameworkCode is the resolver for the createFrameworkCode field.
func (r *mutationResolver) CreateFrameworkCode(ctx context.Context, input api.NewFrameworkCode) (string, error) {
	panic(fmt.Errorf("not implemented: CreateFrameworkCode - createFrameworkCode"))
}

// CreatePatternCode is the resolver for the createPatternCode field.
func (r *mutationResolver) CreatePatternCode(ctx context.Context, input api.NewPatternCode) (string, error) {
	panic(fmt.Errorf("not implemented: CreatePatternCode - createPatternCode"))
}

// CreateAlgorithmCode is the resolver for the createAlgorithmCode field.
func (r *mutationResolver) CreateAlgorithmCode(ctx context.Context, input api.NewAlgorithmCode) (string, error) {
	panic(fmt.Errorf("not implemented: CreateAlgorithmCode - createAlgorithmCode"))
}

// CreateLanguageCode is the resolver for the createLanguageCode field.
func (r *mutationResolver) CreateLanguageCode(ctx context.Context, input api.NewLanguageCode) (string, error) {
	panic(fmt.Errorf("not implemented: CreateLanguageCode - createLanguageCode"))
}

// Mutation returns api.MutationResolver implementation.
func (r *Resolver) Mutation() api.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
