package interceptor

import (
	"context"
	"fmt"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/auth/service"
	"github.com/Shreya1812/ben-and-jerrys/internal/configs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
)

type authInterceptorImpl struct {
	excludedEndpoints map[Endpoint]Endpoint
	authService       service.AuthService
}

func New(excludedEndpoints map[Endpoint]Endpoint, c *configs.Config) (AuthInterceptor, error) {
	authService, err := service.New(c)

	if err != nil {
		return nil, err
	}

	return &authInterceptorImpl{
		excludedEndpoints: excludedEndpoints,
		authService:       authService,
	}, nil
}

func (a authInterceptorImpl) Unary() grpc.UnaryServerInterceptor {
	log.Println(">>>>> Auth Interceptor")

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx, err := a.authorize(ctx, Endpoint(info.FullMethod))

		if err != nil {
			fmt.Printf("Authorization error : %+v", err)
			return nil, err
		}
		return handler(ctx, req)
	}
}

func (a authInterceptorImpl) authorize(ctx context.Context, endpoint Endpoint) (context.Context, error) {
	log.Println(">>>>> Checking Authorization")
	_, ok := a.excludedEndpoints[endpoint]
	if ok {
		//No need of Auth
		log.Printf("Authorization exempted for %+v", endpoint)
		return ctx, nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.Errorf(codes.Unauthenticated, "metadata not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return ctx, status.Errorf(codes.Unauthenticated, "authorization token not provided")
	}

	accessToken := values[0]
	claims, err := a.authService.Verify(ctx, service.Token(accessToken))
	if err != nil {
		return ctx, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	// Saving the email of current logged in user into context.
	// Used to authorize actions on User data.
	// Users cannot delete or update other users data.
	ctx = context.WithValue(ctx, "currentUserEmail", claims.Email)

	log.Println(">>>>> Authorization Completed")
	return ctx, nil
}

func (a authInterceptorImpl) Close() error {
	return a.authService.Close()
}
