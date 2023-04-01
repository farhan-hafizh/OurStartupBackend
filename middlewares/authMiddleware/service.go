package authMiddleware

import (
	"errors"
	"ourstartup/helper"
	"ourstartup/services/user"

	"github.com/golang-jwt/jwt"
)

type Service interface {
	GenerateToken(user user.User) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	jwtSecreteKey string
	encryptionKey string
}

func CreateService(jwtSecreteKey string, encryptionKey string) *jwtService {
	return &jwtService{jwtSecreteKey, encryptionKey}
}

func (s *jwtService) GenerateToken(user user.User) (string, error) {
	// create claim object
	claim := jwt.MapClaims{}
	// initiate claim
	claim["userId"] = user.Id
	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// sign token
	signedToken, err := token.SignedString([]byte(s.jwtSecreteKey))

	if err != nil {
		return signedToken, err
	}
	base64Token, _ := helper.Encrypt([]byte(signedToken), []byte(s.encryptionKey))

	return string(base64Token), nil
}

func (s *jwtService) ValidateToken(paramToken string) (*jwt.Token, error) {
	token, err := helper.Decrypt(paramToken, []byte(s.encryptionKey))
	// parse a token with key function that return any type(interface{}) and error
	decodedToken, err := jwt.Parse(string(token), func(t *jwt.Token) (interface{}, error) {
		// check if the token is signed with HMAC method
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		// if not return the invalid token error
		if !ok {
			return nil, errors.New("Invalid token")
		}
		// if ok, retun the byte so it can be used in jwt.Parse
		return []byte(s.jwtSecreteKey), nil
	})

	if err != nil {
		return decodedToken, err
	}

	return decodedToken, nil
}
