package main

import (
	"context"
	"fmt"
	"log-ingestion/elasticsearch"
	natsjs "log-ingestion/nats-js"
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

	// getting env urls
	esUrl := os.Getenv("ES_URL")
	if esUrl == "" {
		fmt.Println("[ERROR] ES URL not found")
		return
	}

	natsUrl := os.Getenv("NATS_URL")
	if natsUrl == "" {
		fmt.Println("[ERROR] NATS URL not found")
		return
	}

	// ES initialization
	_, err = elasticsearch.New(elasticsearch.Config{
		Addresses: []string{esUrl},
		Username:  "",
		Password:  "",
		CloudId:   "",
	})
	if err != nil {
		fmt.Println("[ERROR] ES connection failed")
		panic(err)
	}

	// creating ILM policy, creating data stream and index template
	if err = elasticsearch.Get().SetupES(ctx); err != nil {
		fmt.Println("[ERROR] ES setup failed")
		panic(err)
	}

	// server
	if err = server.InitializeGrpcServer(ctx); err != nil {
		fmt.Println("[ERROR] Server initialization failed")
		panic(err)
	}

	// connect to nats
	nc, err := natsjs.Connect(natsUrl)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// connect to js
	jc, err := natsjs.JSConnect(nc)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// setting up data streams
	if err = natsjs.InitDataStream(ctx, jc); err != nil {
		fmt.Println(err)
		panic(err)
	}

	sigC := <-sigCh
	fmt.Printf("[MAIN] Terminating signal recieved: %s\n", sigC)
	fmt.Println("[MAIN] Shutting down...")

	cancel()

	fmt.Println("[MAIN] Successfully terminated")
}
