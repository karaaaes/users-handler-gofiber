package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = "SECRET_KEY"

func GenerateJWT(claims *jwt.MapClaims) (string, error) {
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

// VerifikasiJWT digunakan untuk memverifikasi token JWT
func VerifyJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// Verifikasi token
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return token, nil
	}

	return nil, fmt.Errorf("Invalid token")
}

func DecodeToken(token_string string) (jwt.MapClaims, error) {
	token, err := VerifyJWT(token_string)
	if err != nil {
		return nil, err
	}

	claims, isOk := token.Claims.(jwt.MapClaims)
	if isOk && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("Invalid Token")
}
