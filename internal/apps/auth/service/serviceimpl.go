package service

import (
	"context"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/auth/service/model"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/user/service"
	user "github.com/Shreya1812/ben-and-jerrys/internal/apps/user/service/model"
	"github.com/Shreya1812/ben-and-jerrys/internal/commons"
	"github.com/Shreya1812/ben-and-jerrys/internal/configs"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/xerrors"
	"time"
)

type authServiceImpl struct {
	us                   service.UserService
	jwtSecret            string
	jwtExpirationMinutes int32
}

type UserClaims struct {
	Email string
	jwt.StandardClaims
}

func New(c *configs.Config) (AuthService, error) {
	us, err := service.New(c)

	if err != nil {
		return nil, err
	}

	return &authServiceImpl{
		us:                   us,
		jwtSecret:            c.JWTConfig.JWTSecret,
		jwtExpirationMinutes: int32(c.JWTConfig.JwtExpirationMinutes),
	}, nil
}

func (a authServiceImpl) generateJWTToken(email string) (string, error) {
	expirationTime := time.Now().Add(time.Duration(a.jwtExpirationMinutes) * time.Minute)

	claims := &UserClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString([]byte(a.jwtSecret))

	return jwtToken, err
}

func (a authServiceImpl) Login(ctx context.Context, m *model.User) (Token, error) {
	err := a.us.VerifyUser(ctx, &user.User{
		Email:    m.Email,
		Password: m.Password,
	})

	if err != nil {
		return "", err
	}

	t, err := a.generateJWTToken(m.Email)

	return Token(t), err
}

func (a authServiceImpl) Verify(ctx context.Context, token Token) (*UserClaims, error) {
	t, err := jwt.ParseWithClaims(
		string(token),
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, commons.ErrInvalidTokenSigningMethod
			}

			return []byte(a.jwtSecret), nil
		},
	)
	if err != nil {
		return nil, commons.ErrInvalidToken
	}

	claims, ok := t.Claims.(*UserClaims)
	if !ok {
		return nil, commons.ErrInvalidToken
	}

	err = a.us.IsUser(ctx, claims.Email)

	if err != nil {
		if xerrors.Is(err, commons.ErrItemNotFound) {
			return nil, commons.ErrNoSuchUser
		}

		return nil, err
	}
	return claims, nil
}

func (a authServiceImpl) Close() error {
	return a.us.Close()
}
