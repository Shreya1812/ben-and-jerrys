package main

import (
	"context"
	"github.com/Shreya1812/ben-and-jerrys/cmd/client/interceptor"
	auth_pb "github.com/Shreya1812/ben-and-jerrys/internal/proto/auth"
	icecream_pb "github.com/Shreya1812/ben-and-jerrys/internal/proto/icecream"
	user_pb "github.com/Shreya1812/ben-and-jerrys/internal/proto/user"
	"google.golang.org/grpc"
	"log"
)

var user1 = &user_pb.User{
	Email:    "test1@test.com",
	Password: "test1@123",
}
var user2 = &user_pb.User{
	Email:    "test2@test.com",
	Password: "test2@123",
}

func main() {
	unauthorizedCC, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer unauthorizedCC.Close()
	iceCreamClient(unauthorizedCC)
	user1Client(unauthorizedCC)
	user2Client(unauthorizedCC)

	// Creating users
	createUsers(unauthorizedCC)

	log.Println("Logging in Test User 1")
	token1, err := login(unauthorizedCC, user1)
	if err != nil {
		log.Println("Unauthorized", err)
	}
	log.Println("Logged in as Test User 1")

	authorizedCC1, err := newAuthorizedConn(token1)
	iceCreamClient(authorizedCC1)
	user2Client(authorizedCC1)
	user1Client(authorizedCC1)

	log.Println("Logging in Test User 2")
	token2, err := login(unauthorizedCC, user2)
	if err != nil {
		log.Println("Unauthorized", err)
	}
	log.Println("Logged in as Test User 2")

	authorizedCC2, err := newAuthorizedConn(token2)
	iceCreamClient(authorizedCC2)
	user1Client(authorizedCC2)
	user2Client(authorizedCC2)
}

func newAuthorizedConn(token string) (*grpc.ClientConn, error) {
	i := interceptor.New(token)
	return grpc.Dial("localhost:9000", grpc.WithInsecure(), grpc.WithUnaryInterceptor(i.Unary()))
}

func login(cc *grpc.ClientConn, user *user_pb.User) (string, error) {

	client := auth_pb.NewAuthApiClient(cc)
	log.Println("Login:")

	res1, err := client.Login(context.Background(), &auth_pb.LoginRequest{
		User: &auth_pb.User{
			Email:    user.Email,
			Password: user.Password,
		},
	})

	if err != nil {
		return "", err
	}
	return res1.Token, err
}

func createUsers(cc *grpc.ClientConn) {
	client := user_pb.NewUserApiClient(cc)

	log.Println("Creating User Test 1:")
	res1, err := client.Create(context.Background(), &user_pb.CreateRequest{
		User: user1,
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(res1)

	log.Println("Creating User Test 2:")
	res2, err := client.Create(context.Background(), &user_pb.CreateRequest{
		User: user2,
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(res2)
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
	res3, err := client.GetByProductId(context.Background(), &icecream_pb.GetByProductIdRequest{
		ProductId: "10",
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(res3)

	log.Println("Deleting IceCream:")
	res4, err := client.DeleteByProductId(context.Background(), &icecream_pb.DeleteByProductIdRequest{
		ProductId: "10",
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(res4)

	log.Println("Getting Deleted IceCream:")
	res5, err := client.GetByProductId(context.Background(), &icecream_pb.GetByProductIdRequest{
		ProductId: "10",
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(res5)

	log.Println("Getting List of IceCreams")
	res6, err := client.GetList(context.Background(), &icecream_pb.ListRequest{
		PaginationContext: "",
		Limit:             2,
	})
	if err != nil {
		log.Println(err)
	} else {
		log.Println(res6)
		log.Println("Getting Next List of IceCreams")
		res7, err := client.GetList(context.Background(), &icecream_pb.ListRequest{
			PaginationContext: res6.PaginationContext,
			Limit:             2,
		})
		if err != nil {
			log.Println(err)
		}
		log.Println(res7)
	}
}

func user1Client(cc *grpc.ClientConn) {
	client := user_pb.NewUserApiClient(cc)

	log.Println("Updating Test User 1:")
	res1, err := client.Update(context.Background(), &user_pb.UpdateRequest{
		User: &user_pb.User{
			Email:    user1.Email,
			Password: "updatedtest1@123",
		},
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(res1)

	log.Println("Deleting Test User 1:")
	res2, err := client.Delete(context.Background(), &user_pb.DeleteRequest{
		Email: user1.Email,
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(res2)
}

func user2Client(cc *grpc.ClientConn) {
	client := user_pb.NewUserApiClient(cc)

	log.Println("Updating Test User 2:")
	res3, err := client.Update(context.Background(), &user_pb.UpdateRequest{
		User: &user_pb.User{
			Email:    user2.Email,
			Password: "updatedtest2@123",
		},
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(res3)

	log.Println("Deleting Test User 2:")
	res4, err := client.Delete(context.Background(), &user_pb.DeleteRequest{
		Email: user2.Email,
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(res4)
}
