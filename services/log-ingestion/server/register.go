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
	bind4LogChan := make(chan models.Bind4Record, 1000)
	isssLogChan := make(chan models.ISSSRecord, 1000)
	execLogChan := make(chan models.ExecRecord, 1000)
	execveLogChan := make(chan models.ExecveRecord, 1000)
	fanotifyLogChan := make(chan models.FanotifyRecord, 1000)

	// error chans
	dbusErrorChan := make(chan error, 1000)
	connect4ErrorChan := make(chan error, 1000)
	bind4ErrorChan := make(chan error, 1000)
	isssErrorChan := make(chan error, 1000)
	execErrorChan := make(chan error, 1000)
	execveErrorChan := make(chan error, 1000)
	fanotifyErrorChan := make(chan error, 1000)

	// ==========================================================================

	// wg
	var wg sync.WaitGroup

	// register handlers
	pb.RegisterDbusUbitUploaderServer(grpcServer, &handlers.DbusUnitUploaderServer{
		Stream: dbusLogChan,
	})
	pb.RegisterNetworkLogUploaderServer(grpcServer, &handlers.NetworkUploaderServer{
		Connect4Stream: connect4LogChan,
		ISSSStream:     isssLogChan,
		Bind4Stream:    bind4LogChan,
	})
	pb.RegisterExecLogUploaderServer(grpcServer, &handlers.ExecUploaderServer{
		Stream: execLogChan,
	})
	pb.RegisterExecveLogUploaderServer(grpcServer, &handlers.ExecveUploaderServer{
		Stream: execveLogChan,
	})
	pb.RegisterFanotifyUploaderServer(grpcServer, &handlers.FanotifyUploaderServer{
		Stream: fanotifyLogChan,
	})

	// ==========================================================================

	type streamerFunc func()

	// goroutines
	streams := []streamerFunc{
		func() { streamer.DbusStreamer(ctx, dbusLogChan, dbusErrorChan) },
		func() { streamer.Connect4Streamer(ctx, connect4LogChan, connect4ErrorChan) },
		func() { streamer.Bind4Streamer(ctx, bind4LogChan, bind4ErrorChan) },
		func() { streamer.ISSSStreamer(ctx, isssLogChan, isssErrorChan) },
		func() { streamer.ExecStreamer(ctx, execLogChan, execErrorChan) },
		func() { streamer.ExecveStreamer(ctx, execveLogChan, execveErrorChan) },
		func() { streamer.FanotifyStreamer(ctx, fanotifyLogChan, fanotifyErrorChan) },
	}

	for _, stream := range streams {
		wg.Add(1)
		go func() {
			defer wg.Done()
			stream()
		}()
	}

	// ==========================================================================

	// closing
	go func() {
		wg.Wait()
		close(dbusLogChan)
		close(dbusErrorChan)
		close(connect4LogChan)
		close(connect4ErrorChan)
		close(bind4LogChan)
		close(bind4ErrorChan)
		close(isssLogChan)
		close(isssErrorChan)
		close(fanotifyLogChan)
		close(fanotifyErrorChan)
		close(execLogChan)
		close(execErrorChan)
		close(execveLogChan)
		close(execveErrorChan)
	}()
}
