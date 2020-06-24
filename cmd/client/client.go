package main

import (
	"context"
	icecream_pb "github.com/Shreya1812/ben-and-jerrys/internal/proto/icecream"
	user_pb "github.com/Shreya1812/ben-and-jerrys/internal/proto/user"
	"google.golang.org/grpc"
	"log"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	iceCreamClient(cc)
	userClient(cc)
}

func iceCreamClient(cc *grpc.ClientConn) {
	client := icecream_pb.NewIceCreamApiClient(cc)

	log.Println("Creating IceCream:")
	res1, err := client.Create(context.Background(), &icecream_pb.CreateRequest{
		IceCream: &icecream_pb.IceCream{
			ProductId:             "10",
			Name:                  "Icecream Test",
			ImageClosed:           "image/close",
			ImageOpen:             "image/open",
			Description:           "Test Icecream",
			Story:                 "Story",
			SourcingValues:        []string{"source1, source2"},
			Ingredients:           []string{"egg, milk"},
			AllergyInfo:           "no allergy",
			DietaryCertifications: "none",
		},
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(res1)

	log.Println("Updating IceCream:")
	res2, err := client.Update(context.Background(), &icecream_pb.UpdateRequest{
		IceCream: &icecream_pb.IceCream{
			ProductId:             "10",
			Name:                  "Icecream Test Updated",
			ImageClosed:           "image/close/update",
			ImageOpen:             "image/open/update",
			Description:           "Test Icecream Updates",
			Story:                 "Story Updated",
			SourcingValues:        []string{"source1, source2, source3, source4"},
			Ingredients:           []string{"egg, milk, raisins"},
			AllergyInfo:           "no allergy updated",
			DietaryCertifications: "none at all",
		},
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(res2)

	log.Println("Getting IceCream:")
	res3, err := client.GetById(context.Background(), &icecream_pb.GetByIdRequest{
		Id: "10",
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(res3)

	log.Println("Deleting IceCream:")
	res4, err := client.Delete(context.Background(), &icecream_pb.DeleteRequest{
		Id: "10",
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(res4)

	log.Println("Getting Deleted IceCream:")
	res5, err := client.GetById(context.Background(), &icecream_pb.GetByIdRequest{
		Id: "10",
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(res5)
}

func userClient(cc *grpc.ClientConn) {
	client := user_pb.NewUserApiClient(cc)

	log.Println("Creating User:")
	res1, err := client.Create(context.Background(), &user_pb.CreateRequest{
		User: &user_pb.User{
			Email:    "test@test.com",
			Password: "password",
		},
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(res1)

	log.Println("Updating User:")
	res2, err := client.Update(context.Background(), &user_pb.UpdateRequest{
		User: &user_pb.User{
			Email:    "test@test.com",
			Password: "password",
		},
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(res2)

	//log.Println("Deleting User:")
	//res3, err := client.Delete(context.Background(), &user_pb.DeleteRequest{
	//	Email: "test@test.com",
	//})
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(res3)
}
