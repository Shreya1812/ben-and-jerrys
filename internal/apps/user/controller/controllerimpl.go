package controller

import (
	"context"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/user/service"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/user/service/convertor"
	"github.com/Shreya1812/ben-and-jerrys/internal/commons"
	"github.com/Shreya1812/ben-and-jerrys/internal/configs"
	"github.com/Shreya1812/ben-and-jerrys/internal/proto/user"
	"log"
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
	log.Println(">>>>> Creating User")
	err := u.s.CreateUser(ctx, convertor.PbToModel(request.User))

	if err != nil {
		log.Printf("Error creating user %+v", err)
		return nil, commons.GetErrorWithStatus(err)
	}
	log.Println(">>>>> User Created")
	return &user_pb.CreateResponse{}, nil
}

func (u userControllerImpl) Update(ctx context.Context, request *user_pb.UpdateRequest) (*user_pb.UpdateResponse, error) {
	log.Println(">>>>> Updating User")
	currentUserEmail := ctx.Value("currentUserEmail")

	// Users cannot delete or update other users data.
	if currentUserEmail != request.User.Email {
		err := commons.GetErrorWithStatus(commons.ErrPermissionDenied)
		log.Printf("Error updating user %+v", err)
		return nil, err
	}

	err := u.s.UpdateUser(ctx, convertor.PbToModel(request.GetUser()))

	if err != nil {
		log.Printf("Error updating user %+v", err)
		return nil, commons.GetErrorWithStatus(err)
	}

	log.Println(">>>>> User Updated")
	return &user_pb.UpdateResponse{}, nil
}

func (u userControllerImpl) Delete(ctx context.Context, request *user_pb.DeleteRequest) (*user_pb.DeleteResponse, error) {
	log.Println(">>>>> Deleting User")
	currentUserEmail := ctx.Value("currentUserEmail")

	// Users cannot delete or update other users data.
	if currentUserEmail != request.Email {
		err := commons.GetErrorWithStatus(commons.ErrPermissionDenied)
		log.Printf("Error deleting user %+v", err)
		return nil, err
	}

	err := u.s.DeleteUserByEmail(ctx, request.GetEmail())

	if err != nil {
		log.Printf("Error deleting user %+v", err)
		return nil, commons.GetErrorWithStatus(err)
	}

	log.Println(">>>>> User Deleted")
	return &user_pb.DeleteResponse{}, nil
}

func (u userControllerImpl) Close() error {
	return u.s.Close()
}
