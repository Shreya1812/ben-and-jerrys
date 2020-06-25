package service

import (
	"context"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/auth/service/model"
	"io"
)

type Token string

type AuthService interface {
	// Login user using user credentials.
	// Return JWT Token if login is successful
	// or else returns error
	Login(ctx context.Context, m *model.User) (Token, error)

	// Verifies the JWT Token.
	// Returns UserClaims obtained from the JWT Token if verification is successful
	// or else returns error
	Verify(ctx context.Context, token Token) (*UserClaims, error)
	io.Closer
}
