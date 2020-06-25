package interceptor

import (
	"google.golang.org/grpc"
	"io"
)

type Endpoint string

type AuthInterceptor interface {
	Unary() grpc.UnaryServerInterceptor
	io.Closer
}
