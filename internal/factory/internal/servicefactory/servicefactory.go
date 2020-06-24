package servicefactory

import (
	icecream "github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service"
	user "github.com/Shreya1812/ben-and-jerrys/internal/apps/user/service"
	"github.com/Shreya1812/ben-and-jerrys/internal/configs"
)

const (
	host         = "localhost"
	port         = "27017"
	databaseName = "ben-and-jerry"
)

type ServiceFactory struct {
	icecream.IceCreamService
	user.UserService
}

func InitServiceFactory() *ServiceFactory {
	c := &configs.Config{
		MongoDBConfig: &configs.MongoDBConfig{
			Host:         host,
			Port:         port,
			DatabaseName: databaseName,
		},
	}

	return &ServiceFactory{
		IceCreamService: initIceCreamService(c),
		UserService:     initUserService(c),
	}
}

func initIceCreamService(c *configs.Config) icecream.IceCreamService {
	return icecream.New(c)
}

func initUserService(c *configs.Config) user.UserService {
	return user.New(c)
}
