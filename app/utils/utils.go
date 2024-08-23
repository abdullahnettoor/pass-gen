package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/abdullahnettoor/pass-gen/app/config"
	e "github.com/abdullahnettoor/pass-gen/app/models/errors"
	"github.com/abdullahnettoor/pass-gen/app/models/req"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var cfg config.Config

func LoadConfig(config *config.Config) {
	cfg = *config
}

func CreateToken(userID, secretKey string) (string, error) {
	secretKeyByte := []byte(secretKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userID": userID,
			"exp":    time.Now().Add(time.Hour * 240).Unix(),
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
		return "", e.ErrInvalidToken
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

func ValidateToken() (string, error) {
	var token req.Token

	// get user home dir
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// read data from config
	filePath := filepath.Join(homeDir, cfg.ConfigFilePath)
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("--- file-path", filePath)
		fmt.Println("--- ERROR:", err.Error())
		return "", err
	}

	//verify token
	json.Unmarshal(data, &token)

	userID, err := VerifyToken(token.Token, cfg.JwtSecret)
	if err != nil {
		return "", err
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
