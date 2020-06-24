package service

import (
	"context"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/user/service/model"
	"io"
)

type UserService interface {
	CreateUser(ctx context.Context, m *model.User) error
	UpdateUser(ctx context.Context, m *model.User) error
	VerifyUser(ctx context.Context, m *model.User) error
	DeleteUserByEmail(ctx context.Context, email string) error
	IsUser(ctx context.Context, email string) error
	io.Closer
}
