package auth

import (
	"fmt"
	"time"

	"github.com/OmarAouini/employee-manager/config"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

func GenerateToken() (*string, error) {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "golang-employee-api",
		Subject:   "user",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.CONFIG.JwtSecret))
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return &tokenString, nil
}

func ValidateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(config.CONFIG.JwtSecret), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["iat"], claims["iss"])
	} else {
		return fmt.Errorf("unauthorized, %s", err)
	}
	return nil
}

func GetClaimsFromToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(config.CONFIG.JwtSecret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("error getting token claims, %s", err)
	}
	return token.Claims, nil
}
