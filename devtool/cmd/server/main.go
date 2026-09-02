package main

import (
	"log"
	"net"

	"GRPC/devtool/internal/db"
	"GRPC/devtool/internal/repository"
	"GRPC/devtool/internal/service"
	userv1 "GRPC/devtool/pkg/pb/user/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	dsn := "postgres://postgres:postgres@localhost:5432/mydb?sslmode=disable"

	gormDB, err := db.NewPostgresDB(dsn)
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	repo := repository.NewUserRepository(gormDB)
	userServer := service.NewUserServer(repo)

	grpcServer := grpc.NewServer()
	userv1.RegisterUserServiceServer(grpcServer, userServer)

	reflection.Register(grpcServer) // only for testing

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
