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

type NetworkUploaderServer struct {
	pb.UnimplementedNetworkLogUploaderServer
	ISSSStream     chan<- models.ISSSRecord
	Connect4Stream chan<- models.Connect4Record
	Bind4Stream    chan<- models.Bind4Record
}

func (s *NetworkUploaderServer) ISSSUpload(stream pb.NetworkLogUploader_ISSSUploadServer) error {
	var totalChunks int64

	for {
		chunk, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.ISSSAck{
				TotalChunks: totalChunks,
				Message:     "Logs have been received",
			})
		}

		if err != nil {
			fmt.Println("[ERROR] Error receving ISSS logs: %w", err)
			return status.Error(codes.Unknown, "[ISSS] Error occured while getting log")
		}

		s.ISSSStream <- models.ISSSRecord{
			Comm:      chunk.Comm,
			Timestamp: utils.TimestampToUTC(chunk.Timestamp.Seconds, int64(chunk.Timestamp.Nanos)),
			AgentId:   chunk.AgentId,
			OldState:  chunk.OldState,
			NewState:  chunk.NewState,
			SPort:     chunk.SPort,
			DPort:     chunk.DPort,
			Family:    chunk.Family,
			Protocol:  chunk.Protocol,
			SAddr:     chunk.SAddr,
			DAddr:     chunk.DAddr,
			Pid:       chunk.Pid,
			Name:      chunk.Name,
			PPid:      chunk.ParentProcess,
		}
	}
}

func (s *NetworkUploaderServer) Connect4Upload(stream pb.NetworkLogUploader_Connect4UploadServer) error {
	var totalChunks int64

	for {
		chunk, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.Connect4Ack{
				TotalChunks: totalChunks,
				Message:     "Logs have been received",
			})
		}

		if err != nil {
			fmt.Println("[ERROR] Error receving Connect4 logs: %w", err)
			return status.Error(codes.Unknown, "[CONNECT4] Error occured while getting log")
		}

		s.Connect4Stream <- models.Connect4Record{
			UserFamily: chunk.UserFamily,
			UserIPv4:   chunk.UserIPv4,
			UserPort:   chunk.UserPort,
			Family:     chunk.Family,
			Type:       chunk.Type,
			Protocol:   chunk.Protocol,
			Uid:        chunk.Uid,
			Pid:        chunk.Pid,
			Name:       chunk.Name,
			Comm:       chunk.Comm,
			PPid:       chunk.ParentProcess,
			Timestamp:  utils.TimestampToUTC(chunk.Timestamp.Seconds, int64(chunk.Timestamp.Nanos)),
			AgentId:    chunk.AgentId,
		}
	}
}

func (s *NetworkUploaderServer) Bind4Upload(stream pb.NetworkLogUploader_Bind4UploadServer) error {
	var totalChunks int64

	for {
		chunk, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.Bind4Ack{
				TotalChunks: totalChunks,
				Message:     "Logs have been received",
			})
		}

		if err != nil {
			fmt.Println("[ERROR] Error receving Bind4 logs: %w", err)
			return status.Error(codes.Unknown, "[BIND4] Error occured while getting log")
		}

		s.Bind4Stream <- models.Bind4Record{
			UserFamily: chunk.UserFamily,
			UserIPv4:   chunk.UserIPv4,
			UserPort:   chunk.UserPort,
			Family:     chunk.Family,
			Type:       chunk.Type,
			Protocol:   chunk.Protocol,
			Uid:        chunk.Uid,
			Pid:        chunk.Pid,
			Name:       chunk.Name,
			Comm:       chunk.Comm,
			PPid:       chunk.ParentProcess,
			Timestamp:  utils.TimestampToUTC(chunk.Timestamp.Seconds, int64(chunk.Timestamp.Nanos)),
			AgentId:    chunk.AgentId,
		}
	}
}
