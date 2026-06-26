package handlers

import (
	"fmt"
	"io"
	"log-ingestion/models"
	"log-ingestion/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ExecveUploaderServer struct {
	pb.UnimplementedExecveLogUploaderServer
	Stream chan<- models.ExecveRecord
}

func (s *ExecveUploaderServer) ExecveUpload(stream pb.ExecveLogUploader_ExecveUploadServer) error {
	var totalChunks int64

	for {
		chunk, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.ExecveAck{
				TotalChunks: totalChunks,
				Message:     "Logs have been received",
			})
		}

		if err != nil {
			fmt.Println("[ERROR] Error receving execve logs: %w", err)
			return status.Error(codes.Unknown, "[EXECVE] Error occured while getting log")
		}

		s.Stream <- models.ExecveRecord{
			Pid:       uint32(chunk.Pid),
			Name:      chunk.Name,
			Comm:      chunk.Comm,
			PPid:      chunk.ParentProcess,
			Timestamp: chunk.Timestamp,
			AgentId:   chunk.AgentId,
		}
	}
}
