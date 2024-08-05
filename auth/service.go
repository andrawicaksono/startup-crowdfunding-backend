package auth

import "github.com/golang-jwt/jwt/v5"

type Service interface {
	GenerateToken(id int) (string, error)
}

type service struct {
}

func NewService() *service {
	return &service{}
}

var SECRET_KEY = []byte("BWASTARTUP_s3C23t_K3y")

func (s *service) GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
