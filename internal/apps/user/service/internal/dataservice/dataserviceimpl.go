package dataservice

import (
	"context"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/user/service/internal/config"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/user/service/internal/internal/convertor"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/user/service/internal/internal/storage/db"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/user/service/model"
)

type userDataServiceImpl struct {
	db db.UserDB
}

func New(config *config.UserConfig) (UserDataService, error) {
	database, err := db.New(config.MongoDBConfig)

	if err != nil {
		return nil, err
	}
	return &userDataServiceImpl{
		db: database,
	}, nil
}

func (u userDataServiceImpl) Create(ctx context.Context, m *model.User) error {
	return u.db.Create(ctx, convertor.ModelToData(m))
}

func (u userDataServiceImpl) Update(ctx context.Context, m *model.User) error {
	return u.db.Update(ctx, convertor.ModelToData(m))
}

func (u userDataServiceImpl) DeleteByEmail(ctx context.Context, email string) error {
	return u.db.Delete(ctx, email)
}

func (u userDataServiceImpl) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	d, err := u.db.GetByEmail(ctx, email)

	if err != nil {
		return nil, err
	}
	return convertor.DataToModel(d), nil
}

func (u userDataServiceImpl) Close() error {
	return u.db.Close()
}
