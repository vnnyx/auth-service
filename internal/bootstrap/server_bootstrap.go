package bootstrap

import (
	"log"
	"net"

	"github.com/vnnyx/auth-service/internal/delivery/grpc/di"
	"github.com/vnnyx/auth-service/internal/infrastructure"
	pb "github.com/vnnyx/auth-service/pb/auth"
	"google.golang.org/grpc"
)

func StartServer() {
	cfg := infrastructure.NewConfig()
	grpcDelivery := di.InitializeGRPCAAuthServer()
	srv := grpc.NewServer()
	pb.RegisterAuthServiceServer(srv, grpcDelivery)

	log.Println("Starting RPC server at", cfg.ServiceAuthPort)

	l, err := net.Listen("tcp", cfg.ServiceAuthPort)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", cfg.ServiceAuthPort, err)
	}

	log.Fatal(srv.Serve(l))
}
