package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func CreateToken(userID, secretKey string) (string, error) {
	secretKeyByte := []byte(secretKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userID": userID,
			"exp":    time.Now().Add(time.Hour * 10).Unix(),
		})

	tokenString, err := token.SignedString(secretKeyByte)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString, secretKey string) (string, error) {
	secretKeyByte := []byte(secretKey)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKeyByte, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", err
	}

	userID, ok := claims["userID"].(string)
	if !ok {
		return "", fmt.Errorf("token claims is not matching")
	}

	return userID, nil
}

func HashPassword(password string) (string, error) {
	p, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashedPassword := string(p)
	return hashedPassword, err
}

func CompareHashedPassword(dbPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
}
