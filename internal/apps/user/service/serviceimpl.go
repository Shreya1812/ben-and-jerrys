package service

import (
	"context"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/user/service/internal/config"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/user/service/internal/dataservice"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/user/service/model"
	"github.com/Shreya1812/ben-and-jerrys/internal/configs"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type userServiceImpl struct {
	ds dataservice.UserDataService
}

func New(c *configs.Config) UserService {
	return &userServiceImpl{
		ds: dataservice.New(config.GetUserConfig(c)),
	}
}

type Claims struct {
	Email string
	jwt.StandardClaims
}

var JwtSecret = []byte("ben_and_jerry_jwt_secret")

const JwtExpirationMinutes = 5

func generateJWTToken(email string) (string, error) {
	expirationTime := time.Now().Add(JwtExpirationMinutes * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	jwtToken, err := token.SignedString(JwtSecret)

	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func (u userServiceImpl) CreateUser(ctx context.Context, m *model.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	m.Password = string(hashed)
	return u.ds.Create(ctx, m)
}

func (u userServiceImpl) UpdateUser(ctx context.Context, m *model.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	m.Password = string(hashed)
	return u.ds.Update(ctx, m)
}

func (u userServiceImpl) VerifyUser(ctx context.Context, m *model.User) error {
	user, err := u.ds.GetByEmail(ctx, m.Email)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(m.Password))
	if err != nil {
		return err // Auth error
	}
	return nil
}

func (u userServiceImpl) DeleteUserByEmail(ctx context.Context, email string) error {
	return u.ds.DeleteByEmail(ctx, email)
}

func (u userServiceImpl) IsUser(ctx context.Context, email string) error {
	_, err := u.ds.GetByEmail(ctx, email)
	return err
}

func (u userServiceImpl) Close() error {
	return u.ds.Close()
}
