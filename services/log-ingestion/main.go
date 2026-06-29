package main

import (
	"context"
	"fmt"
	"log-ingestion/elasticsearch"
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

	// ES initialization
	_, err = elasticsearch.New(elasticsearch.Config{
		Addresses: []string{"http://es-dev:9200"},
		Username:  "",
		Password:  "",
		CloudId:   "",
	})
	if err != nil {
		fmt.Println("[ERROR] ES connection failed")
		panic(err)
	}

	if err = elasticsearch.Get().SetupES(ctx); err != nil {
		fmt.Println("[ERROR] ES setup failed")
		panic(err)
	}

	// server
	if err = server.InitializeGrpcServer(ctx); err != nil {
		fmt.Println("[ERROR] Server initialization failed")
		panic(err)
	}

	sigC := <-sigCh
	fmt.Printf("[MAIN] Terminating signal recieved: %s\n", sigC)
	fmt.Println("[MAIN] Shutting down...")

	cancel()

	fmt.Println("[MAIN] Successfully terminated")
}
