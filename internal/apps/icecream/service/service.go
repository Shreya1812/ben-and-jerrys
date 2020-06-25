package service

import (
	"context"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/model"
	"io"
)

type IceCreamService interface {
	CreateIceCream(ctx context.Context, m *model.IceCream) (*model.IceCream, error)
	UpdateIceCream(ctx context.Context, m *model.IceCream) (*model.IceCream, error)
	DeleteIceCreamByProductId(ctx context.Context, pId string) (*model.IceCream, error)
	GetIceCreamByProductId(ctx context.Context, pId string) (*model.IceCream, error)
	GetIceCreamList(ctx context.Context, searchOptions *model.IceCreamSearchOptions) (*model.IceCreamListResult, error)
	io.Closer
}
