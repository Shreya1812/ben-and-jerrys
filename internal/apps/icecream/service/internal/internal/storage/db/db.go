package db

import (
	"context"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/internal/internal/data"
	"io"
)

type IceCreamDB interface {
	Create(ctx context.Context, d *data.IceCream) (*data.IceCream, error)
	Update(ctx context.Context, d *data.IceCream) (*data.IceCream, error)
	Delete(ctx context.Context, pId string) (*data.IceCream, error)
	GetById(ctx context.Context, pId string) (*data.IceCream, error)
	io.Closer
}
