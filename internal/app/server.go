package app

import (
	"github.com/DrusGalkin/Auth-gRPC/internal/server"
	auth "github.com/DrusGalkin/Auth-gRPC/pkg/api/g_rpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	auth.RegisterAuthServiceServer(s, &server.AuthServer{})

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
