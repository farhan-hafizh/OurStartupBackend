package auth

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

type Service interface {
	GenerateToken(userId int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
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

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	// parse a token with key function that return any type(interface{}) and error
	decodedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		// check if the token is signed with HMAC method
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		// if not return the invalid token error
		if !ok {
			return nil, errors.New("Invalid token")
		}
		// if ok, retun the byte so it can be used in jwt.Parse
		return []byte(SECRETE_KEY), nil
	})

	if err != nil {
		return decodedToken, err
	}

	return decodedToken, nil
}
