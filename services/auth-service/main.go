package main

import (
	"auth-service/db"
	"auth-service/server"
	"context"
)

func main() {
	ctx := context.Background()

	// connect db
	_, err := db.DbConnect(ctx)
	if err != nil {
		panic(err)
	}

	// seed
	err = db.InitializeDBSeed(ctx)
	if err != nil {
		panic(err)
	}

	// start server
	err = server.InitializeServer()
	if err != nil {
		panic(err)
	}
}
