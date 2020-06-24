package service

import (
	"context"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/model"
	"io"
)

type IceCreamService interface {
	CreateIceCream(ctx context.Context, m *model.IceCream) (*model.IceCream, error)
	UpdateIceCream(ctx context.Context, m *model.IceCream) (*model.IceCream, error)
	DeleteIceCreamById(ctx context.Context, pId string) (*model.IceCream, error)
	GetIceCreamById(ctx context.Context, pId string) (*model.IceCream, error)
	io.Closer
}
