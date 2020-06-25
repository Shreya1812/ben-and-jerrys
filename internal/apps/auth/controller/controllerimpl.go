package controller

import (
	"context"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/auth/service"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/auth/service/convertor"
	"github.com/Shreya1812/ben-and-jerrys/internal/commons"
	"github.com/Shreya1812/ben-and-jerrys/internal/configs"
	auth_pb "github.com/Shreya1812/ben-and-jerrys/internal/proto/auth"
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
	t, err := a.s.Login(ctx, convertor.PbToModel(request.User))

	if err != nil {
		return nil, commons.GetErrorWithStatus(err)
	}

	return &auth_pb.LoginResponse{Token: string(t)}, nil
}

func (a authControllerImpl) Close() error {
	return a.s.Close()
}
