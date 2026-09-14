package main

import (
	"log"
	"net"

	handler "GRPC/grpc-1/services/orders/handler/orders"
	"GRPC/grpc-1/services/orders/service"

	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	ordersService := service.NewOrderService()
	handler.NewGrpcOrdersService(grpcServer, ordersService)

	log.Println("Starting gRPC server on", s.addr)

	return grpcServer.Serve(lis)
}
