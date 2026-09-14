package resolver

import (
	"context"
	"strconv"

	"GRPC/graphql-2/graph/model"
	"GRPC/graphql-2/models"
	"GRPC/graphql-2/repository"
)

func CreateUser(ctx context.Context, name, email string) (*model.User, error) {
	user := models.User{
		Name:  name,
		Email: email,
	}

	err := repository.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:    strconv.Itoa(int(user.ID)),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
