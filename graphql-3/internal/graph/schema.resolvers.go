package graph

// THIS CODE WILL BE UPDATED WITH SCHEMA CHANGES. PREVIOUS IMPLEMENTATION FOR SCHEMA CHANGES WILL BE KEPT IN THE COMMENT SECTION. IMPLEMENTATION FOR UNCHANGED SCHEMA WILL BE KEPT.

import (
	"GRPC/graphql-3/internal/graph/model"
	"context"
)

type Resolver struct{}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, name string, email string) (*model.User, error) {
	u, err := r.UserSvc.CreateUser(ctx, name, email)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, name *string, email *string) (*model.User, error) {
	u, err := r.UserSvc.UpdateUser(ctx, id, name, email)
	if err != nil {
		return nil, err
	}
	return &model.User{ID: u.ID, Name: u.Name, Email: u.Email}, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	return r.UserSvc.DeleteUser(ctx, id)
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	u, err := r.UserSvc.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return &model.User{ID: u.ID, Name: u.Name, Email: u.Email}, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := r.UserSvc.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.User, 0, len(users))
	for _, u := range users {
		result = append(result, &model.User{ID: u.ID, Name: u.Name, Email: u.Email})
	}
	return result, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type (
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
)

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	type Resolver struct{}
*/
