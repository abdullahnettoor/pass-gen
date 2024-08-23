package usecase

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/abdullahnettoor/pass-gen/app/config"
	"github.com/abdullahnettoor/pass-gen/app/models/req"
	"github.com/abdullahnettoor/pass-gen/app/models/res"
	"github.com/abdullahnettoor/pass-gen/app/repo"
	"github.com/abdullahnettoor/pass-gen/app/utils"
)

var configData *config.Config

func LoadConfig(credential *config.Config) {
	configData = credential
}

func Signup(user *req.User) (*res.User, error) {
	var err error

	if user.Password != user.ConfirmPassword {
		log.Fatal("password and confirm password not matching")
	}

	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	result, err := repo.Signup(user)
	if err != nil {
		return nil, err
	}

	jwtToken, err := utils.CreateToken(result.UserID, configData.JwtSecret)
	if err != nil {
		return nil, err
	}

	tokenModel := req.Token{Token: jwtToken}
	byteTokenModel, err := json.Marshal(tokenModel)
	if err != nil {
		return nil, err
	}

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configDir := filepath.Join(userHomeDir, configData.ConfigPath)
	confFileDir := filepath.Join(userHomeDir, configData.ConfigFilePath)

	err = os.MkdirAll(configDir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	file, err := os.Create(confFileDir)
	if err != nil {
		fmt.Println("----", err, confFileDir)
		return nil, err
	}

	_, err = file.Write(byteTokenModel)
	if err != nil {
		return nil, err
	}

	return result, nil
}
