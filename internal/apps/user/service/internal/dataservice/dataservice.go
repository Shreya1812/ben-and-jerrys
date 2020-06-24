package dataservice

import (
	"context"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/user/service/model"
	"io"
)

type UserDataService interface {
	Create(ctx context.Context, m *model.User) error
	Update(ctx context.Context, m *model.User) error
	DeleteByEmail(ctx context.Context, email string) error
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	io.Closer
}
