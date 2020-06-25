package dataservice

import (
	"context"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/model"
	"io"
)

//go:generate mockgen -destination=mock_dataservice.go -package=dataservice -source=dataservice.go
type IceCreamDataService interface {
	Create(ctx context.Context, m *model.IceCream) (*model.IceCream, error)
	Update(ctx context.Context, m *model.IceCream) (*model.IceCream, error)
	DeleteByProductId(ctx context.Context, pId string) (*model.IceCream, error)
	GetByProductId(ctx context.Context, pId string) (*model.IceCream, error)
	GetList(ctx context.Context, options *model.IceCreamSearchOptions) (*model.IceCreamListResult, error)
	io.Closer
}
