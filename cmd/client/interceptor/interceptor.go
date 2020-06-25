package interceptor

import (
	"google.golang.org/grpc"
	"io"
)

type AuthInterceptor interface {
	Unary() grpc.UnaryClientInterceptor
	io.Closer
}
