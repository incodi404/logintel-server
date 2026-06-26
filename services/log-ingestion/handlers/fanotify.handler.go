package handlers

import (
	"fmt"
	"io"
	"log-ingestion/models"
	"log-ingestion/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FanotifyUploaderServer struct {
	pb.UnimplementedFanotifyUploaderServer
	Stream chan<- models.FanotifyRecord
}

func (s *FanotifyUploaderServer) FanotifyUpload(stream pb.FanotifyUploader_FanotifyUploadServer) error {
	var totalChunks int64

	for {
		chunk, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.FanotifyAck{
				TotalChunks: totalChunks,
				Message:     "Logs have been received",
			})
		}

		if err != nil {
			fmt.Println("[ERROR] Error receving fanotify logs: %w", err)
			return status.Error(codes.Unknown, "[FANOTIFY] Error occured while getting log")
		}

		s.Stream <- models.FanotifyRecord{
			Comm:      chunk.Comm,
			Timestamp: chunk.Timestamp,
			AgentId:   chunk.AgentId,
			Pid:       chunk.Pid,
			Name:      chunk.Name,
			PPid:      chunk.ParentProcess,
			Path:      chunk.Path,
			Events:    chunk.Events,
		}
	}
}
