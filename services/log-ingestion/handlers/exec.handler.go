package handlers

import (
	"fmt"
	"io"
	"log-ingestion/models"
	"log-ingestion/pb"
	"log-ingestion/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ExecUploaderServer struct {
	pb.UnimplementedExecLogUploaderServer
	Stream chan<- models.ExecRecord
}

func (s *ExecUploaderServer) ExecUpload(stream pb.ExecLogUploader_ExecUploadServer) error {
	var totalChunks int64

	for {
		chunk, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.ExecAck{
				TotalChunks: totalChunks,
				Message:     "Logs have been received",
			})
		}

		if err != nil {
			fmt.Println("[ERROR] Error receving exec logs: %w", err)
			return status.Error(codes.Unknown, "[DBUS] Error occured while getting log")
		}

		s.Stream <- models.ExecRecord{
			Filename:  chunk.Filename,
			Pid:       chunk.Pid,
			OldPid:    chunk.OldPid,
			Uid:       chunk.Uid,
			Comm:      chunk.Comm,
			Timestamp: utils.TimestampToUTC(chunk.Timestamp.Seconds, int64(chunk.Timestamp.Nanos)),
			AgentId:   chunk.AgentId,
		}
	}
}
