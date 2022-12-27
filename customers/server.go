package customers

import (
	"log"
	"net"

	pbCustomer "github.com/kaualimamartins/gRPC-Go-Gorm/customers/proto/pb"
	"google.golang.org/grpc"
)

func CreateGRPCServer() {
	listener, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatalf("failed to listen 8000 port: %v", err)
	}

	opts := []grpc.ServerOption{}

	grpcServer := grpc.NewServer(opts...)

	customerServer := NewCustomerServer()

	pbCustomer.RegisterCustomerServiceServer(grpcServer, customerServer)

	log.Println("Listening on port 8000")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
