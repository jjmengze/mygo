package graph

import "github.com/jjmengze/mygo/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userService usecase.UserService
}

func NewResolver(userService usecase.UserService) *Resolver {
	return &Resolver{userService: userService}
}
