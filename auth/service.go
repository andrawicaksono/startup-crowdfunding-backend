package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GenerateToken(id int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type service struct {
}

func NewService() *service {
	return &service{}
}

var SECRET_KEY = []byte("BWASTARTUP_s3C23t_K3y")
var LOGIN_EXPIRATION_DURATION = time.Duration(15) * time.Minute

func (s *service) GenerateToken(userID int) (string, error) {
	type MyClaims struct {
		jwt.RegisteredClaims
		UserID int `json:"user_id"`
	}

	claims := MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(LOGIN_EXPIRATION_DURATION)),
		},
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(SECRET_KEY)
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

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
