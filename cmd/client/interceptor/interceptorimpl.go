package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type authInterceptorImpl struct {
	token string
}

func New(token string) AuthInterceptor {
	return &authInterceptorImpl{
		token: token,
	}
}

func (a *authInterceptorImpl) Unary() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{},
		cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		return invoker(a.attachToken(ctx), method, req, reply, cc, opts...)
	}
}

func (a *authInterceptorImpl) attachToken(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", a.token)
}

func (a *authInterceptorImpl) Close() error {
	panic("implement me")
}
