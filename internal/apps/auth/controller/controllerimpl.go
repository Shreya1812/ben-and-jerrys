package controller

import (
	"context"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/auth/service"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/auth/service/convertor"
	"github.com/Shreya1812/ben-and-jerrys/internal/commons"
	"github.com/Shreya1812/ben-and-jerrys/internal/configs"
	auth_pb "github.com/Shreya1812/ben-and-jerrys/internal/proto/auth"
	"log"
)

type authControllerImpl struct {
	s service.AuthService
}

func New(c *configs.Config) (AuthController, error) {
	s, err := service.New(c)

	if err != nil {
		return nil, err
	}
	return &authControllerImpl{
		s: s,
	}, nil
}

func (a authControllerImpl) Login(ctx context.Context, request *auth_pb.LoginRequest) (*auth_pb.LoginResponse, error) {
	log.Println(">>>>> User Logging In")
	t, err := a.s.Login(ctx, convertor.PbToModel(request.User))

	if err != nil {
		log.Printf("Error loggin in user: %+v", err)
		return nil, commons.GetErrorWithStatus(err)
	}

	log.Println(">>>>> User Logged In")
	return &auth_pb.LoginResponse{Token: string(t)}, nil
}

func (a authControllerImpl) Close() error {
	return a.s.Close()
}
