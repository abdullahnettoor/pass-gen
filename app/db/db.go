package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PassGenUser struct {
	ID       int    `gorm:"primaryKey,AUTO_INCREMENT"`
	UserName string `gorm:"unique"`
	Password string
}

type PassGenPasswordStore struct {
	ID       int
	UserID   int
	Name     string
	Password string
}

func InitDB(dbconnection string) (*gorm.DB, error) {
	DB, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbconnection,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Println("ERROR: DB OPEN:", err.Error())
		return nil, err
	}

	err = DB.AutoMigrate(&PassGenUser{}, &PassGenPasswordStore{})
	if err != nil {
		log.Println("ERROR: DB MIGRATION:", err.Error())
		return nil, err
	}

	return DB, nil
}
