package service

import (
	"context"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/internal/config"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/internal/dataservice"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/model"
	"github.com/Shreya1812/ben-and-jerrys/internal/commons"
	"github.com/Shreya1812/ben-and-jerrys/internal/configs"
	"gopkg.in/go-playground/validator.v9"
)

type iceCreamServiceImpl struct {
	ds dataservice.IceCreamDataService
	v  *validator.Validate
}

func New(c *configs.Config) (IceCreamService, error) {
	ds, err := dataservice.New(config.GetIceCreamConfig(c))

	if err != nil {
		return nil, err
	}

	return &iceCreamServiceImpl{
		ds: ds,
		v:  model.GetIceCreamValidator(),
	}, nil
}

func (i *iceCreamServiceImpl) validateIceCream(m *model.IceCream) error {
	err := i.v.Struct(m)
	if err != nil {
		return commons.ErrValidation{Err: err}
	}
	return nil
}

func (i *iceCreamServiceImpl) CreateIceCream(ctx context.Context, m *model.IceCream) (*model.IceCream, error) {
	if err := i.validateIceCream(m); err != nil {
		return nil, err
	}
	return i.ds.Create(ctx, m)
}

func (i *iceCreamServiceImpl) UpdateIceCream(ctx context.Context, m *model.IceCream) (*model.IceCream, error) {
	if err := i.validateIceCream(m); err != nil {
		return nil, err
	}
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
