package resolvers

import "NintendoCenter/gateway/internal/protos"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	client protos.GameCollectionClient
}

func NewResolver(client protos.GameCollectionClient) *Resolver {
	return &Resolver{client: client}
}
