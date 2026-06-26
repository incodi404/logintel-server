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

	// error chans
	dbusErrorChan := make(chan error, 1000)

	// wg
	var wg sync.WaitGroup

	// register
	pb.RegisterDbusUbitUploaderServer(grpcServer, &handlers.DbusUnitUploaderServer{
		Stream: dbusLogChan,
	})

	// goroutines
	wg.Add(1)
	go func() {
		defer wg.Done()
		streamer.DbusStreamer(ctx, dbusLogChan, dbusErrorChan)
	}()

	// closing
	go func() {
		wg.Wait()
		close(dbusLogChan)
		close(dbusErrorChan)
	}()

}
