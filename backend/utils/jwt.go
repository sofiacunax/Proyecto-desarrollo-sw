package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("mi_clave_super_secreta")

func GenerateToken(
	id int,
	email string,
	rol string,
) (string, error) {

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":    id,
			"email": email,
			"rol":   rol,
			"exp":   time.Now().Add(24 * time.Hour).Unix(),
		},
	)

	return token.SignedString(SecretKey)
}
