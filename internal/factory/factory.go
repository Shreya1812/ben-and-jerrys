package factory

import (
	auth "github.com/Shreya1812/ben-and-jerrys/internal/apps/auth/controller"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/auth/interceptor"
	icecream "github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/controller"
	user "github.com/Shreya1812/ben-and-jerrys/internal/apps/user/controller"
	"github.com/Shreya1812/ben-and-jerrys/internal/configs"
	"github.com/pkg/errors"
	"os"
	"strconv"
)

var excludedEndPoints = map[interceptor.Endpoint]interceptor.Endpoint{
	interceptor.Endpoint("/user.UserApi/Create"): interceptor.Endpoint("/user.UserApi/Create"),
	interceptor.Endpoint("/auth.AuthApi/Login"):  interceptor.Endpoint("/auth.AuthApi/Login"),
}

type Factory struct {
	iceCreamController icecream.IceCreamController
	userController     user.UserController
	authController     auth.AuthController
	authInterceptor    interceptor.AuthInterceptor
}

func (c *Factory) GetIceCreamController() icecream.IceCreamController {
	return c.iceCreamController
}

func (c *Factory) GetUserController() user.UserController {
	return c.userController
}

func (c *Factory) GetAuthController() auth.AuthController {
	return c.authController
}

func (c *Factory) GetAuthInterceptor() interceptor.AuthInterceptor {
	return c.authInterceptor
}

func InitFactory() (*Factory, error) {

	c, err := getConfig()

	if err != nil {
		return nil, errors.Wrap(err, "configuration error")
	}

	iceCreamController, err := icecream.New(c)
	if err != nil {
		return nil, errors.Wrap(err, "error initializing iceCreamController")
	}

	userController, err := user.New(c)
	if err != nil {
		return nil, errors.Wrap(err, "error initializing userController")
	}

	authController, err := auth.New(c)
	if err != nil {
		return nil, errors.Wrap(err, "error initializing authController")
	}

	authInterceptor, err := interceptor.New(excludedEndPoints, c)
	if err != nil {
		return nil, errors.Wrap(err, "error initializing authInterceptor")
	}

	return &Factory{
		iceCreamController: iceCreamController,
		userController:     userController,
		authController:     authController,
		authInterceptor:    authInterceptor,
	}, nil
}

func (c *Factory) DisposeController() error {
	if err := c.iceCreamController.Close(); err != nil {
		return err
	}

	if err := c.userController.Close(); err != nil {
		return err
	}

	if err := c.authController.Close(); err != nil {
		return err
	}

	if err := c.authInterceptor.Close(); err != nil {
		return err
	}

	return nil
}

func getConfig() (*configs.Config, error) {

	mongoHost := os.Getenv("MONGO_HOST")
	if len(mongoHost) == 0 {
		mongoHost = "localhost"
	}

	mongoPort := os.Getenv("MONGO_PORT")
	if len(mongoPort) == 0 {
		mongoPort = "27017"
	}

	mongoDBName := os.Getenv("MONGO_DB_NAME")
	if len(mongoDBName) == 0 {
		mongoDBName = "ben-and-jerrys"
	}

	jWTSecret := os.Getenv("JWT_SECRET")
	if len(jWTSecret) == 0 {
		jWTSecret = "JWT_SECRET"
	}

	jwtExpirationMinutesStr := os.Getenv("JWT_EXPIRY_MIN")
	if len(jwtExpirationMinutesStr) == 0 {
		jwtExpirationMinutesStr = "5"
	}
	jwtExpirationMinutes, err := strconv.Atoi(jwtExpirationMinutesStr)

	if err != nil {
		return nil, err
	}

	c := &configs.Config{
		MongoDBConfig: &configs.MongoDBConfig{
			Host:         mongoHost,
			Port:         mongoPort,
			DatabaseName: mongoDBName,
		},
		JWTConfig: &configs.JWTConfig{
			JWTSecret:            jWTSecret,
			JwtExpirationMinutes: jwtExpirationMinutes,
		},
	}

	return c, nil
}
