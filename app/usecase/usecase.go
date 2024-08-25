package usecase

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/abdullahnettoor/pass-gen/app/config"
	e "github.com/abdullahnettoor/pass-gen/app/models/errors"
	"github.com/abdullahnettoor/pass-gen/app/models/req"
	"github.com/abdullahnettoor/pass-gen/app/models/res"
	"github.com/abdullahnettoor/pass-gen/app/pkg/generator"
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
		fmt.Println(user.Password, "--", user.ConfirmPassword)
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

	jwtToken, err := utils.CreateToken(result.ID, configData.JwtSecret)
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

func Login(user *req.User) error {
	user.UserName = strings.ToLower(user.UserName)

	userData, err := repo.Login(user)
	if err != nil {
		return err
	}

	err = utils.CompareHashedPassword(userData.Password, user.Password)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		return e.ErrInvalidPassword
	}
	fmt.Println("User Data", userData)

	//create jwt token
	jwtToken, err := utils.CreateToken(userData.ID, configData.JwtSecret)
	if err != nil {
		return err
	}

	tokenModel := req.Token{Token: jwtToken}
	byteTokenModel, err := json.Marshal(tokenModel)
	if err != nil {
		return err
	}

	// store in config file
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configDir := filepath.Join(userHomeDir, configData.ConfigPath)
	confFileDir := filepath.Join(configDir, configData.ConfigFilePath)

	// configDir := filepath.Join(userHomeDir, ".pass_gen")
	// confFileDir := filepath.Join(configDir, "config.json")

	fmt.Println("Config directory:", configDir)
	fmt.Println("Config file path:", confFileDir)

	// Create the directory if it doesn't exist
	err = os.MkdirAll(configDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}
	fmt.Println("Directory created successfully")

	fmt.Println("Attempting to create file at:", confFileDir)
	file, err := os.Create(confFileDir)
	if err != nil {
		fmt.Println("hello", err)
		return fmt.Errorf("failed to create config file: %w", err)
	}
	defer file.Close()

	// Write to the file
	_, err = file.Write(byteTokenModel)
	if err != nil {
		return fmt.Errorf("failed to write to config file: %w", err)
	}

	fmt.Println("File created and written successfully")
	return nil
	// err = os.MkdirAll(configDir, os.ModePerm)
	// if err != nil {
	// 	return err
	// }

	// file, err := os.Create(confFileDir)
	// if err != nil {
	// 	fmt.Println("hello")
	// 	return err
	// }
	// defer file.Close()

	// _, err = file.Write(byteTokenModel)
	// if err != nil {
	// 	fmt.Println("how r u")
	// 	return err
	// }

	// return nil
}

func StoreSecret(credential *req.Credential) (err error) {
	credential.CipherText, err = generator.Erncypt([]byte(configData.CipherSecret), []byte(credential.Secret))
	if err != nil {
		return err
	}

	return repo.StoreSecret(credential)
}
