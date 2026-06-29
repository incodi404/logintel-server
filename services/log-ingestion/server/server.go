package server

import (
	"context"
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"
)

func InitializeGrpcServer(ctx context.Context) error {
	port := os.Getenv("PORT")
	fmt.Println("ENV PORT: ", port)
	if port == "" {
		port = "5051"
	}

	list, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		return fmt.Errorf("[SERVER] Error occured while listening: %w", err)
	}

	grpcServer := grpc.NewServer()

	// register
	RegisterHandlers(grpcServer, ctx)

	fmt.Println("[INFO] Server is listening on 5051 abcdwdc")
	if err = grpcServer.Serve(list); err != nil {
		return fmt.Errorf("[SERVER] Error serving gRPC server: %w", err)
	}

	return nil
}
