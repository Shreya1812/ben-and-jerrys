package db

import (
	"context"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/user/service/internal/internal/data"
	"io"
)

type UserDB interface {
	Create(ctx context.Context, d *data.User) error
	Update(ctx context.Context, d *data.User) error
	Delete(ctx context.Context, email string) error
	GetByEmail(ctx context.Context, email string) (*data.User, error)
	io.Closer
}
