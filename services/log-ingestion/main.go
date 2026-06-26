package main

import (
	"context"
	"fmt"
	"log-ingestion/server"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	// sig
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// ctx
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("[ERROR] Error loading .env: %w", err)
		return
	}

	// server
	server.InitializeGrpcServer(ctx)

	sigC := <-sigCh
	fmt.Printf("[MAIN] Terminating signal recieved: %s\n", sigC)
	fmt.Println("[MAIN] Shutting down...")

	cancel()

	fmt.Println("[MAIN] Successfully terminated")
}
