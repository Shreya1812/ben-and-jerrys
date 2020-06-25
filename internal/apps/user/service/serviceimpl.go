package service

import (
	"context"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/user/service/internal/config"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/user/service/internal/dataservice"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/user/service/model"
	"github.com/Shreya1812/ben-and-jerrys/internal/commons"
	"github.com/Shreya1812/ben-and-jerrys/internal/configs"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
)

type userServiceImpl struct {
	ds dataservice.UserDataService
	v  *validator.Validate
}

func New(c *configs.Config) (UserService, error) {
	ds, err := dataservice.New(config.GetUserConfig(c))

	if err != nil {
		return nil, err
	}

	return &userServiceImpl{
		ds: ds,
		v:  model.GetUserValidator(),
	}, nil
}

func (u *userServiceImpl) hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (u *userServiceImpl) validateUser(m *model.User) error {
	err := u.v.Struct(m)
	if err != nil {
		return commons.ErrValidation{Err: err}
	}
	return nil
}

func (u *userServiceImpl) CreateUser(ctx context.Context, m *model.User) error {
	if err := u.validateUser(m); err != nil {
		return err
	}

	hashed, err := u.hashPassword(m.Password)
	if err != nil {
		return err
	}
	m.Password = string(hashed)
	return u.ds.Create(ctx, m)
}

func (u *userServiceImpl) UpdateUser(ctx context.Context, m *model.User) error {
	if err := u.validateUser(m); err != nil {
		return err
	}
	hashed, err := u.hashPassword(m.Password)
	if err != nil {
		return err
	}
	m.Password = string(hashed)
	return u.ds.Update(ctx, m)
}

func (u *userServiceImpl) VerifyUser(ctx context.Context, m *model.User) error {
	user, err := u.ds.GetByEmail(ctx, m.Email)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(m.Password))
	if err != nil {
		return commons.ErrAuthentication
	}
	return nil
}

func (u *userServiceImpl) DeleteUserByEmail(ctx context.Context, email string) error {
	return u.ds.DeleteByEmail(ctx, email)
}

func (u *userServiceImpl) IsUser(ctx context.Context, email string) error {
	_, err := u.ds.GetByEmail(ctx, email)
	return err
}

func (u *userServiceImpl) Close() error {
	return u.ds.Close()
}
