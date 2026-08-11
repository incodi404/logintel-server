package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	natsjs "rule-engine/nats-js"
	"rule-engine/pool"
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
	natsUrl := os.Getenv("NATS_URL")
	if natsUrl == "" {
		fmt.Println("[ERROR] NATS URL not found")
		return
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
	if err = natsjs.InitConsumer(ctx, jc); err != nil {
		fmt.Println(err)
		panic(err)
	}

	// start worker pools
	pool.StartWorkerPool(ctx)

	sigC := <-sigCh
	fmt.Printf("[MAIN] Terminating signal recieved: %s\n", sigC)
	fmt.Println("[MAIN] Shutting down...")

	cancel()

	fmt.Println("[MAIN] Successfully terminated")
}
