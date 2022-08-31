package helper

import (
	"api_gateway/models"
	"errors"
	"github.com/golang-jwt/jwt"
)

var jwtSecret = "4c73fc2d42887a9ba1678384b611900254a3eeee2a2eeeba57a11eb17c45c47d"

func SignJwt(user *models.User) (string, error) {

	sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{
		"id":   user.Id,
		"role": user.Role,
	})

	token, err := sign.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseJwt(token string) (jwt.MapClaims, error) {
	data, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != jwt.SigningMethodHS256 {
			return nil, errors.New("invalid token")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := data.Claims.(jwt.MapClaims)
	if !ok || !data.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
