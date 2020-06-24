package controller

import (
	"github.com/Shreya1812/ben-and-jerrys/internal/proto/icecream"
	"io"
)

type IceCreamController interface {
	icecream_pb.IceCreamApiServer
	io.Closer
}
