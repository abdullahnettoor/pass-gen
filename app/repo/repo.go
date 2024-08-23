package repo

import (
	e "github.com/abdullahnettoor/pass-gen/app/models/errors"
	"github.com/abdullahnettoor/pass-gen/app/models/req"
	"github.com/abdullahnettoor/pass-gen/app/models/res"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitRepository(DBs *gorm.DB) {
	db = DBs
}

func Signup(user *req.User) (*res.User, error) {
	var UserRes res.User
	query := "INSERT INTO safex_users (user_name, password) SELECT $1, $2 WHERE NOT EXISTS (SELECT 1 FROM safex_users WHERE user_name = $1) RETURNING *"
	result := db.Raw(query, user.UserName, user.Password).Scan(&UserRes)
	if result.Error != nil {
		return nil, e.ErrDb
	}

	if result.RowsAffected == 0 {
		return nil, e.ErrConflict
	}
	return &UserRes, nil
}
