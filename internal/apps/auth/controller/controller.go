package controller

import (
	auth_pb "github.com/Shreya1812/ben-and-jerrys/internal/proto/auth"
	"io"
)

type AuthController interface {
	auth_pb.AuthApiServer
	io.Closer
}
