package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/caiostarke/apiAndGRPC/logger_service/logger_pb"
	"github.com/caiostarke/apiAndGRPC/logger_service/platform/database"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	logger_pb.UnimplementedLoggerServiceServer
}

func RunGRPCServer() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	logger_pb.RegisterLoggerServiceServer(s, &Server{})
	reflection.Register(s)

	go func() {
		fmt.Println("Starting server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("\nStopping the server")
	s.Stop()
	fmt.Println("Closing the listener")
	lis.Close()
	fmt.Println("End of program. =(")
}

func (*Server) CreateLog(ctx context.Context, req *logger_pb.CreateLogRequest) (*logger_pb.CreateLogResponse, error) {
	fmt.Println("Creating a log")

	db, err := database.OpenDBConnection()
	if err != nil {
		return &logger_pb.CreateLogResponse{}, err
	}

	defer db.DB.Disconnect(context.TODO())

	log := req.GetLog()

	if err = db.CreateLog(log); err != nil {
		return &logger_pb.CreateLogResponse{}, err
	}

	return &logger_pb.CreateLogResponse{
		Log: log,
	}, nil
}
