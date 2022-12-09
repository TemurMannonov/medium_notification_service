package main

import (
	"log"
	"net"

	// pb "github.com/TemurMannonov/medium_notification_service/genproto/user_service"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/TemurMannonov/medium_notification_service/config"
	// "github.com/TemurMannonov/medium_notification_service/service"
)

func main() {
	cfg := config.Load(".")

	// userService := service.NewUserService(strg, inMemory)

	lis, err := net.Listen("tcp", cfg.GrpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	// pb.RegisterUserServiceServer(s, userService)

	log.Println("Grpc server started in port ", cfg.GrpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while listening: %v", err)
	}

}
