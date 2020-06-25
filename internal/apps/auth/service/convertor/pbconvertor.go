package convertor

import (
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/auth/service/model"
	"github.com/Shreya1812/ben-and-jerrys/internal/proto/auth"
)

func PbToModel(pb *auth_pb.User) *model.User {
	return &model.User{
		Email:    pb.Email,
		Password: pb.Password,
	}
}
