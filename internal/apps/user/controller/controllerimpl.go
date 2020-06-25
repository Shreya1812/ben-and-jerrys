package controller

import (
	"context"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/user/service"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/user/service/convertor"
	"github.com/Shreya1812/ben-and-jerrys/internal/commons"
	"github.com/Shreya1812/ben-and-jerrys/internal/configs"
	"github.com/Shreya1812/ben-and-jerrys/internal/proto/user"
)

type userControllerImpl struct {
	s service.UserService
}

func New(c *configs.Config) (UserController, error) {
	s, err := service.New(c)

	if err != nil {
		return nil, err
	}
	return &userControllerImpl{
		s: s,
	}, nil
}

func (u userControllerImpl) Create(ctx context.Context, request *user_pb.CreateRequest) (*user_pb.CreateResponse, error) {
	err := u.s.CreateUser(ctx, convertor.PbToModel(request.User))

	if err != nil {
		return nil, commons.GetErrorWithStatus(err)
	}

	return &user_pb.CreateResponse{}, nil
}

func (u userControllerImpl) Update(ctx context.Context, request *user_pb.UpdateRequest) (*user_pb.UpdateResponse, error) {
	currentUserEmail := ctx.Value("currentUserEmail")

	if currentUserEmail != request.User.Email {
		return nil, commons.GetErrorWithStatus(commons.ErrPermissionDenied)
	}

	err := u.s.UpdateUser(ctx, convertor.PbToModel(request.GetUser()))

	if err != nil {
		return nil, commons.GetErrorWithStatus(err)
	}

	return &user_pb.UpdateResponse{}, nil
}

func (u userControllerImpl) Delete(ctx context.Context, request *user_pb.DeleteRequest) (*user_pb.DeleteResponse, error) {
	currentUserEmail := ctx.Value("currentUserEmail")

	if currentUserEmail != request.Email {
		return nil, commons.GetErrorWithStatus(commons.ErrPermissionDenied)
	}

	err := u.s.DeleteUserByEmail(ctx, request.GetEmail())

	if err != nil {
		return nil, commons.GetErrorWithStatus(err)
	}

	return &user_pb.DeleteResponse{}, nil
}

func (u userControllerImpl) Close() error {
	return u.s.Close()
}
