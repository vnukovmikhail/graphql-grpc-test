package resolver

import (
	"context"
	"strconv"

	"GRPC/knight/graph/model"
	"GRPC/knight/repository"
)

func User(ctx context.Context) ([]*model.User, error) {
	user, err := repository.GetUsers()
	if err != nil {
		return nil, err
	}

	var results []*model.User
	for _, u := range user {
		results = append(results, &model.User{
			ID:    strconv.Itoa(int(u.ID)),
			Name:  u.Name,
			Email: u.Email,
		})
	}

	return results, nil
}
