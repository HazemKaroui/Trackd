package main

import (
	"fmt"
	"log"
	"net"
	"tracker/config"
	"tracker/repository"
	"tracker/tracker"
	"tracker/usecase"

	"google.golang.org/grpc"
)

func main() {
	db := config.ConnectDB()
	TodoRepository := repository.InitTrackedRepository(db)
	trackerUsecase := usecase.InitUserUsecase(TodoRepository)

	s := tracker.InitServer(trackerUsecase)

	grpcServer := grpc.NewServer()

	tracker.RegisterTrackerServiceServer(grpcServer, &s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3000))
	if err != nil {
		log.Println("failed to listen: ", err)
	}
	fmt.Println("Listen to port 3000")

	if err := grpcServer.Serve(lis); err != nil {
		log.Println("failed to serve: ", err)
	}
}
