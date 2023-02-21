package auth

import "github.com/golang-jwt/jwt"

type Service interface {
	GenerateToken(userId int) (string, error)
}

type jwtService struct{}

var SECRETE_KEY = []byte("S3CR3T3")

func CreateService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userId int) (string, error) {
	// create claim object
	claim := jwt.MapClaims{}
	// initiate claim
	claim["user_id"] = userId
	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// sign token
	signedToken, err := token.SignedString(SECRETE_KEY)

	if err != nil {
		return signedToken, nil
	}

	return signedToken, nil
}
