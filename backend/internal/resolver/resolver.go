package resolver

import (
	"github.com/tf63/code-api/internal/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Tr repository.TodoRepository
}
