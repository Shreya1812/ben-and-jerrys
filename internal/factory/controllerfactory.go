package factory

import (
	icecream "github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/controller"
	user "github.com/Shreya1812/ben-and-jerrys/internal/apps/user/controller"
	"github.com/Shreya1812/ben-and-jerrys/internal/factory/internal/servicefactory"
)

type ControllerFactory struct {
	iceCreamController icecream.IceCreamController
	userController     user.UserController
}

func (c *ControllerFactory) GetIceCreamController() icecream.IceCreamController {
	return c.iceCreamController
}

func (c *ControllerFactory) GetUserController() user.UserController {
	return c.userController
}

func InitControllerFactory() *ControllerFactory {
	sf := servicefactory.InitServiceFactory()

	return &ControllerFactory{
		iceCreamController: initIceCreamController(sf),
		userController:     initUserController(sf),
	}
}

func (c *ControllerFactory) DisposeController() error {
	if err := c.iceCreamController.Close(); err != nil {
		return err
	}

	if err := c.userController.Close(); err != nil {
		return err
	}

	return nil
}

func initIceCreamController(sf *servicefactory.ServiceFactory) icecream.IceCreamController {
	return icecream.New(sf.IceCreamService)
}

func initUserController(sf *servicefactory.ServiceFactory) user.UserController {
	return user.New(sf.UserService)
}
