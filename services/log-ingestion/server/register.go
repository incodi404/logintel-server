package server

import (
	"context"
	"log-ingestion/handlers"
	"log-ingestion/models"
	"log-ingestion/pb"
	"log-ingestion/streamer"
	"sync"

	"google.golang.org/grpc"
)

func RegisterHandlers(grpcServer *grpc.Server, ctx context.Context) {
	// channels
	dbusLogChan := make(chan models.DbusUnitRecord, 1000)
	connect4LogChan := make(chan models.Connect4Record, 1000)

	// error chans
	dbusErrorChan := make(chan error, 1000)
	connect4ErrorChan := make(chan error, 1000)

	// wg
	var wg sync.WaitGroup

	// register
	pb.RegisterDbusUbitUploaderServer(grpcServer, &handlers.DbusUnitUploaderServer{
		Stream: dbusLogChan,
	})
	pb.RegisterNetworkLogUploaderServer(grpcServer, &handlers.NetworkUploaderServer{
		Connect4Stream: connect4LogChan,
	})

	// goroutines
	wg.Add(2)
	go func() {
		defer wg.Done()
		streamer.DbusStreamer(ctx, dbusLogChan, dbusErrorChan)
	}()

	go func() {
		defer wg.Done()
		streamer.Connect4Streamer(ctx, connect4LogChan, connect4ErrorChan)
	}()

	// closing
	go func() {
		wg.Wait()
		close(dbusLogChan)
		close(dbusErrorChan)
	}()

}
