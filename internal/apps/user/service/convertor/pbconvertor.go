package convertor

import (
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/user/service/model"
	"github.com/Shreya1812/ben-and-jerrys/internal/proto/user"
)

func PbToModel(pb *user_pb.User) *model.User {
	return &model.User{
		Email:    pb.Email,
		Password: pb.Password,
	}
}
