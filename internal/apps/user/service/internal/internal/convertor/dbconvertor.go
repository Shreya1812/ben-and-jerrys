package convertor

import (
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/user/service/internal/internal/data"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/user/service/model"
)

func ModelToData(m *model.User) *data.User {
	return &data.User{
		Email:    m.Email,
		Password: m.Password,
	}
}

func DataToModel(d *data.User) *model.User {
	return &model.User{
		Email:    d.Email,
		Password: d.Password,
	}
}
