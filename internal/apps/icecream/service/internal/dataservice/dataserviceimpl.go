package dataservice

import (
	"context"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/internal/config"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/internal/internal/convertor"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/internal/internal/storage/db"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/model"
)

type iceCreamDataServiceImpl struct {
	db db.IceCreamDB
}

func New(config *config.IceCreamConfig) (IceCreamDataService, error) {
	database, err := db.New(config.MongoDBConfig)

	if err != nil {
		return nil, err
	}
	return &iceCreamDataServiceImpl{
		db: database,
	}, nil
}

func (i *iceCreamDataServiceImpl) Create(ctx context.Context, m *model.IceCream) (*model.IceCream, error) {
	d, err := i.db.Create(ctx, convertor.ModelToData(m))

	if err != nil {
		return nil, err
	}
	return convertor.DataToModel(d), nil
}

func (i *iceCreamDataServiceImpl) Update(ctx context.Context, m *model.IceCream) (*model.IceCream, error) {
	d, err := i.db.Update(ctx, convertor.ModelToData(m))

	if err != nil {
		return nil, err
	}
	return convertor.DataToModel(d), nil
}

func (i *iceCreamDataServiceImpl) DeleteById(ctx context.Context, pId string) (*model.IceCream, error) {
	d, err := i.db.Delete(ctx, pId)

	if err != nil {
		return nil, err
	}
	return convertor.DataToModel(d), nil
}

func (i *iceCreamDataServiceImpl) GetById(ctx context.Context, pId string) (*model.IceCream, error) {
	d, err := i.db.GetById(ctx, pId)

	if err != nil {
		return nil, err
	}

	return convertor.DataToModel(d), nil
}

func (i *iceCreamDataServiceImpl) Close() error {
	return i.db.Close()
}
