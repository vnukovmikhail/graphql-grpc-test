package repository

import (
	"GRPC/graphql-2/config"
	"GRPC/graphql-2/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User

	err := config.DB.Find(&users).Error
	return users, err
}

func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}
