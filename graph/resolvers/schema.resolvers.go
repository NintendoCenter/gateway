package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"NintendoCenter/gateway/graph/generated"
	"NintendoCenter/gateway/graph/model"
	"NintendoCenter/gateway/internal/mapper"
	"NintendoCenter/gateway/internal/protos"
)

func (r *mutationResolver) CreateNotification(ctx context.Context, input model.NotificationInput) (*model.Notification, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveNotification(ctx context.Context, gameID string) ([]*model.Notification, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Find(ctx context.Context, input model.SearchInput) ([]*model.Game, error) {
	request := protos.FindGameRequest{Title: *input.Title}
	response, err := r.client.FindGame(ctx, &request)
	if err != nil {
		return nil, err
	}

	games := mapper.GameMapper{}.GamesToResponse(response.Games)
	return games, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Games(ctx context.Context) ([]*model.Game, error) {
	panic(fmt.Errorf("not implemented"))
}
