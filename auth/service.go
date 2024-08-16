package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type Service interface {
	GenerateToken(id int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type service struct {
	SECRET_KEY                []byte
	LOGIN_EXPIRATION_DURATION time.Duration
}

func NewService(viper *viper.Viper) *service {
	secretKey := viper.GetString("jwt.secret_key")
	loginExpirationDuration := viper.GetInt("jwt.login_expiration_duration")

	SECRET_KEY := []byte(secretKey)
	LOGIN_EXPIRATION_DURATION := time.Duration(loginExpirationDuration) * time.Minute

	return &service{SECRET_KEY, LOGIN_EXPIRATION_DURATION}
}

func (s *service) GenerateToken(userID int) (string, error) {
	type MyClaims struct {
		jwt.RegisteredClaims
		UserID int `json:"user_id"`
	}

	claims := MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.LOGIN_EXPIRATION_DURATION)),
		},
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(s.SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *service) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		method, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		} else if method != jwt.SigningMethodHS256 {
			return nil, errors.New("invalid token")
		}

		return s.SECRET_KEY, nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
