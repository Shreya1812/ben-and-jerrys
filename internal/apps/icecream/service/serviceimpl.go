package service

import (
	"context"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/internal/config"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/internal/dataservice"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/model"
	"github.com/Shreya1812/ben-and-jerrys/internal/configs"
)

type iceCreamServiceImpl struct {
	ds dataservice.IceCreamDataService
}

func New(c *configs.Config) IceCreamService {
	return &iceCreamServiceImpl{
		ds: dataservice.New(config.GetIceCreamConfig(c)),
	}
}

func (i *iceCreamServiceImpl) CreateIceCream(ctx context.Context, m *model.IceCream) (*model.IceCream, error) {
	return i.ds.Create(ctx, m)
}

func (i *iceCreamServiceImpl) UpdateIceCream(ctx context.Context, m *model.IceCream) (*model.IceCream, error) {
	return i.ds.Update(ctx, m)
}

func (i *iceCreamServiceImpl) DeleteIceCreamById(ctx context.Context, pId string) (*model.IceCream, error) {
	return i.ds.DeleteById(ctx, pId)
}

func (i *iceCreamServiceImpl) GetIceCreamById(ctx context.Context, pId string) (*model.IceCream, error) {
	return i.ds.GetById(ctx, pId)
}

func (i *iceCreamServiceImpl) Close() error {
	return i.ds.Close()
}