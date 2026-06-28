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

type DbusUnitUploaderServer struct {
	pb.UnimplementedDbusUbitUploaderServer
	Stream chan<- models.DbusUnitRecord
}

func (s *DbusUnitUploaderServer) DbusUnitUpload(stream pb.DbusUbitUploader_DbusUnitUploadServer) error {
	var totalChunks int64

	for {
		chunk, err := stream.Recv()

		if err == io.EOF {
			uploadResult := &pb.DbusUnitAck{
				TotalChunks: totalChunks,
				Message:     "Logs have been received",
			}

			return stream.SendAndClose(uploadResult)
		}

		if err != nil {
			fmt.Println("[ERROR] Error receving dbus logs: %w", err)
			return status.Error(codes.Unknown, "[DBUS] Error occured while getting log")
		}

		s.Stream <- models.DbusUnitRecord{
			Name:        chunk.Name,
			Description: chunk.Description,
			LoadState:   chunk.LoadState,
			ActiveState: chunk.ActiveState,
			Substate:    chunk.SubState,
			Followed:    chunk.Followed,
			Path:        chunk.Path,
			JobId:       chunk.JobId,
			JobType:     chunk.JobType,
			JobPath:     chunk.JobPath,
			Timestamp:   utils.TimestampToUTC(chunk.Timestamp.Seconds, int64(chunk.Timestamp.Nanos)),
		}
	}
}
