package auth

import "github.com/golang-jwt/jwt/v5"

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("BWASTARTUP_s3cr3t_k3y")

func NewService() *jwtService {
	return &jwtService{}
}

func(s *jwtService) GenerateToken(userID int) (string,error) {
	payload := jwt.MapClaims{}

	payload["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}