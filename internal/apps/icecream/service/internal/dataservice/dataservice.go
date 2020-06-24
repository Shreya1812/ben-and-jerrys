package dataservice

import (
	"context"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/model"
	"io"
)

type IceCreamDataService interface {
	Create(ctx context.Context, m *model.IceCream) (*model.IceCream, error)
	Update(ctx context.Context, m *model.IceCream) (*model.IceCream, error)
	DeleteById(ctx context.Context, pId string) (*model.IceCream, error)
	GetById(ctx context.Context, pId string) (*model.IceCream, error)
	io.Closer
}
