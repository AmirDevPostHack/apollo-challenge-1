package graph

import "github.com/amir/graphql-c1/internal/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	svc *service.Service
}

func NewResolver(svc *service.Service) *Resolver {
	return &Resolver{
		svc: svc,
	}
}
