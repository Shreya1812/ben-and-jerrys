package controller

import (
	"github.com/Shreya1812/ben-and-jerrys/internal/proto/user"
	"io"
)

type UserController interface {
	user_pb.UserApiServer
	io.Closer
}
