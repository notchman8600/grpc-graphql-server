package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"github.com/vektah/gqlparser/v2/formatter"
	"notchman.tech/gateway/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	rand, _ := rand.Int(rand.Reader, big.NewInt(100))
	user := &model.User{
		Name: input.Name,
		ID:   fmt.Sprintf("T%d", rand),
	}
	r.users = append(r.users, user)
	return user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return []*model.User{
		{
			ID:   "user:1",
			Name: "Federation Taro",
		},
		{
			ID:   "user:2",
			Name: "Go Taro",
		},
	}, nil
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
