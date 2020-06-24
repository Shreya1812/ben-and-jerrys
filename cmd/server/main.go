package main

import (
	"github.com/Shreya1812/ben-and-jerrys/internal/factory"
	icecream_pb "github.com/Shreya1812/ben-and-jerrys/internal/proto/icecream"
	user_pb "github.com/Shreya1812/ben-and-jerrys/internal/proto/user"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()

	cf := factory.InitControllerFactory()
	icecream_pb.RegisterIceCreamApiServer(s, cf.GetIceCreamController())
	user_pb.RegisterUserApiServer(s, cf.GetUserController())

	go func() {
		log.Printf("Starting server")

		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for Control C
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	log.Print("Stopping the server")
	s.Stop()
	log.Print("Closing the listener")
	if err := lis.Close(); err != nil {
		log.Print(err)
	}
	log.Printf("Closing all connections")
	if err := cf.DisposeController(); err != nil {
		log.Print(err)
	}
	log.Print("Server shutdown")
}
