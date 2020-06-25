package commons

import (
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrValidation struct {
	Err error
}

func (e ErrValidation) Error() string {
	return fmt.Sprintf("validation error: %v", e.Err)
}

type ErrServerStartup struct {
	Err error
}

func (e ErrServerStartup) Error() string {
	return fmt.Sprintf("server startup error: %v", e.Err)
}

type ErrServerShutdown struct {
	Err error
}

func (e ErrServerShutdown) Error() string {
	return fmt.Sprintf("server shutdown error: %v", e.Err)
}

var (
	ErrItemAlreadyExists = errors.New("item already exists")
	ErrItemNotFound      = errors.New("item not found")

	ErrNoSuchUser       = errors.New("no such user")
	ErrAuthentication   = errors.New("invalid credentials")
	ErrPermissionDenied = errors.New("unauthorized to perform this action")

	ErrInvalidTokenSigningMethod = errors.New("invalid JWT token signing method")
	ErrInvalidToken              = errors.New("invalid token")
)

func getErrorStatus(err error) codes.Code {
	var code codes.Code

	switch err {
	case ErrItemAlreadyExists:
		code = codes.AlreadyExists
	case ErrItemNotFound, ErrNoSuchUser:
		code = codes.NotFound

	case ErrAuthentication:
		code = codes.Unauthenticated
	case ErrPermissionDenied:
		code = codes.PermissionDenied

	case ErrInvalidTokenSigningMethod, ErrInvalidToken:
		code = codes.InvalidArgument
	default:
		code = codes.Internal
	}

	switch err.(type) {
	case ErrValidation:
		code = codes.InvalidArgument
	case ErrServerStartup, ErrServerShutdown:
		code = codes.Internal
	}

	return code
}

func GetErrorWithStatus(err error) error {
	return status.Errorf(
		getErrorStatus(err),
		fmt.Sprintf(err.Error()),
	)
}
