package db

import (
	"context"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/internal/internal/data"
	"io"
)

type IceCreamDB interface {
	Create(ctx context.Context, d *data.IceCream) (*data.IceCream, error)
	Update(ctx context.Context, d *data.IceCream) (*data.IceCream, error)
	DeleteByProductId(ctx context.Context, pId string) (*data.IceCream, error)
	GetByProductId(ctx context.Context, pId string) (*data.IceCream, error)
	GetList(ctx context.Context, options *data.IceCreamSearchOptions) (*data.IceCreamListResult, error)
	io.Closer
}
