package convertor

import (
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/user/service/model"
	"github.com/Shreya1812/ben-and-jerrys/internal/proto/user"
	"strings"
)

func PbToModel(pb *user_pb.User) *model.User {
	return &model.User{
		Email:    strings.Trim(pb.Email, " "),
		Password: strings.Trim(pb.Password, " "),
	}
}
