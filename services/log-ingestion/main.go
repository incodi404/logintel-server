package main

import (
	"fmt"
	"io"
	"net"
	"test_server/pb"

	"google.golang.org/grpc"
)

type execUploaderServer struct {
	pb.UnimplementedExecLogUploaderServer
}

type execveUploaderServer struct {
	pb.UnimplementedExecveLogUploaderServer
}

type networkUploaderServer struct {
	pb.UnimplementedNetworkLogUploaderServer
}

type fanotifyUploaderServer struct {
	pb.UnimplementedFanotifyUploaderServer
}

type dbusUnitUploaderServer struct {
	pb.UnimplementedDbusUbitUploaderServer
}

func (s *execUploaderServer) ExecUpload(stream pb.ExecLogUploader_ExecUploadServer) error {
	var totalChunks int64

	for {
		chunk, err := stream.Recv()

		if err == io.EOF {
			result := &pb.ExecAck{
				Message:     "Logs received successfully",
				TotalChunks: totalChunks,
			}

			return stream.SendAndClose(result)
		}

		if err != nil {
			return fmt.Errorf("Error receiving exec logs: %w", err)
		}

		totalChunks++

		fmt.Println("[EXEC] Received: ", chunk)
	}
}

func (s *execveUploaderServer) ExecveUpload(stream pb.ExecveLogUploader_ExecveUploadServer) error {
	var totalChunks int64

	for {
		chunk, err := stream.Recv()

		if err == io.EOF {
			result := &pb.ExecveAck{
				Message:     "Logs has been reached",
				TotalChunks: totalChunks,
			}

			return stream.SendAndClose(result)
		}

		if err != nil {
			return fmt.Errorf("Error receiving execve logs: %w", err)
		}

		totalChunks++

		fmt.Println("[EXECVE] Received: ", chunk)
	}
}

func (s *networkUploaderServer) ISSSUpload(stream pb.NetworkLogUploader_ISSSUploadServer) error {
	var totalChunks int64

	for {
		chunk, err := stream.Recv()

		if err == io.EOF {
			uploadResult := &pb.ISSSAck{
				Message:     "Logs have been reached",
				TotalChunks: totalChunks,
			}
			return stream.SendAndClose(uploadResult)
		}

		if err != nil {
			return fmt.Errorf("Error receiving ISSS logs: %w", err)
		}

		fmt.Println("[ISSS] Log received: %w", chunk)
	}
}

func (s *networkUploaderServer) Connect4Upload(stream pb.NetworkLogUploader_Connect4UploadServer) error {
	var totalChunks int64

	for {
		chunk, err := stream.Recv()

		if err == io.EOF {
			uploadResult := &pb.Connect4Ack{
				TotalChunks: totalChunks,
				Message:     "Logs have been reached",
			}

			return stream.SendAndClose(uploadResult)
		}

		if err != nil {
			return fmt.Errorf("[CONNECT4] Error occured: %w", err)
		}

		fmt.Println("[CONNECT4] Log received: %w", chunk)
	}
}

func (s *networkUploaderServer) Bind4Upload(stream pb.NetworkLogUploader_Bind4UploadServer) error {
	var totalChunks int64

	for {
		chunk, err := stream.Recv()

		if err == io.EOF {
			uploadResult := &pb.Bind4Ack{
				TotalChunks: totalChunks,
				Message:     "Logs have been reached",
			}

			return stream.SendAndClose(uploadResult)
		}

		if err != nil {
			return fmt.Errorf("[BIND4] Error occured: %w", err)
		}

		fmt.Println("[CONNECT4] Log received: %w", chunk)
	}
}

func (s *fanotifyUploaderServer) FanotifyUpload(stream pb.FanotifyUploader_FanotifyUploadServer) error {
	var totalChunks int64

	for {
		chunk, err := stream.Recv()

		if err == io.EOF {
			uploadResult := &pb.FanotifyAck{
				TotalChunks: totalChunks,
				Message:     "Logs have been reached",
			}

			return stream.SendAndClose(uploadResult)
		}

		if err != nil {
			return fmt.Errorf("[FANOTIFY] Error occured: %w", err)
		}

		fmt.Println("[FANOTIFY] Log received: ", chunk)
	}
}

func (s *dbusUnitUploaderServer) DbusUnitUpload(stream pb.DbusUbitUploader_DbusUnitUploadServer) error {
	var totalChunks int64

	for {
		chunk, err := stream.Recv()

		if err == io.EOF {
			uploadResult := &pb.DbusUnitAck{
				TotalChunks: totalChunks,
				Message:     "Logs have been reached",
			}

			return stream.SendAndClose(uploadResult)
		}

		if err != nil {
			return fmt.Errorf("[DBUS] Error occured: %w", err)
		}

		fmt.Println("[DBUS] Log received: ", chunk)
	}
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:5051")
	if err != nil {
		fmt.Println("[ERROR] Error listening on 5051: %w", err)
		return
	}

	grpcServer := grpc.NewServer()

	pb.RegisterExecLogUploaderServer(grpcServer, &execUploaderServer{})
	pb.RegisterExecveLogUploaderServer(grpcServer, &execveUploaderServer{})
	pb.RegisterNetworkLogUploaderServer(grpcServer, &networkUploaderServer{})
	pb.RegisterFanotifyUploaderServer(grpcServer, &fanotifyUploaderServer{})
	pb.RegisterDbusUbitUploaderServer(grpcServer, &dbusUnitUploaderServer{})

	fmt.Println("[INFO] Server is listening on 5051")
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("[ERROR] Error serrving on 5051: %w", err)
		return
	}
}
