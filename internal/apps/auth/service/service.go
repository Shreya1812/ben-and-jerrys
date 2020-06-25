package service

import (
	"context"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/auth/service/model"
	"io"
)

type Token string

type AuthService interface {
	Login(ctx context.Context, m *model.User) (Token, error)
	Verify(ctx context.Context, token Token) (*UserClaims, error)
	io.Closer
}
