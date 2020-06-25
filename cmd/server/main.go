package main

import (
	"github.com/Shreya1812/ben-and-jerrys/internal/factory"
	auth_pb "github.com/Shreya1812/ben-and-jerrys/internal/proto/auth"
	icecream_pb "github.com/Shreya1812/ben-and-jerrys/internal/proto/icecream"
	user_pb "github.com/Shreya1812/ben-and-jerrys/internal/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:9000")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	cf, err := factory.InitFactory()

	if err != nil {
		log.Fatalf("Failed to initialize factory: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(cf.GetAuthInterceptor().Unary()))
	reflection.Register(s)

	icecream_pb.RegisterIceCreamApiServer(s, cf.GetIceCreamController())
	user_pb.RegisterUserApiServer(s, cf.GetUserController())
	auth_pb.RegisterAuthApiServer(s, cf.GetAuthController())

	go func() {
		log.Println(">>>>> Starting server")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
		log.Println(">>>>> Started server")
	}()

	// Wait for Control C
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	log.Println(">>>>> Stopping server")
	s.Stop()
	log.Println(">>>>> Closing the listener")
	if err := lis.Close(); err != nil {
		log.Print(err)
	}
	log.Println(">>>>> Closing all connections")
	if err := cf.DisposeController(); err != nil {
		log.Print(err)
	}
	log.Println(">>>>> Stopped sever")
}
