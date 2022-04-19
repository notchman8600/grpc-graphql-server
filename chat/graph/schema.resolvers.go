package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strings"

	"github.com/vektah/gqlparser/v2/formatter"
	"notchman.tech/chat/graph/model"
)

func (r *mutationResolver) SendMessage(ctx context.Context, input model.NewMessage) (*model.Chat, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.Chat, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Service(ctx context.Context) (*model.Service, error) {
	s := new(strings.Builder)
	f := formatter.NewFormatter(s)
	// parsedSchema is in the generated code
	f.FormatSchema(parsedSchema)

	service := model.Service{
		Name:    "gqlgen-service",
		Version: "0.1.0",
		Schema:  s.String(),
	}
	return &service, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func (r *mutationResolver) CreateChat(ctx context.Context, input model.NewMessage) (*model.Chat, error) {
	chat := &model.Chat{
		Message: input.Message,
		UserID:  input.UserID,
	}
	r.chats = append(r.chats, chat)
	return chat, nil
}
func (r *queryResolver) Chats(ctx context.Context) ([]*model.Chat, error) {
	return []*model.Chat{
		{
			UserID:  "example-user-id-1",
			Message: "hogehoge",
		},
		{
			UserID:  "example-user-id-2",
			Message: "hogehoge",
		},
	}, nil
}
