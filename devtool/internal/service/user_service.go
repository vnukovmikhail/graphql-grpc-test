package service

import (
	"context"
	"errors"

	"GRPC/devtool/internal/model"
	"GRPC/devtool/internal/repository"
	userv1 "GRPC/devtool/pkg/pb/user/v1"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserServer struct {
	userv1.UnimplementedUserServiceServer
	repo repository.UserRepository
}

func NewUserServer(repo repository.UserRepository) *UserServer {
	return &UserServer{repo: repo}
}

func (s *UserServer) Register(ctx context.Context, req *userv1.RegisterRequest) (*userv1.RegisterResponse, error) {
	if req.GetEmail() == "" || req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "email and password required")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to process password")
	}

	u := &model.User{
		Email:        req.GetEmail(),
		Username:     req.GetUsername(),
		PasswordHash: string(hash),
	}

	if err := s.repo.Create(ctx, u); err != nil {
		return nil, status.Error(codes.AlreadyExists, "user already exists")
	}

	return &userv1.RegisterResponse{
		User: &userv1.User{
			Id:       u.ID,
			Email:    u.Email,
			Username: u.Username,
		},
	}, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {
	u, err := s.repo.GetById(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "database error")
	}

	return &userv1.GetUserResponse{
		User: &userv1.User{
			Id:       u.ID,
			Email:    u.Email,
			Username: u.Username,
		},
	}, nil
}
