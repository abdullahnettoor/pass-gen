package repo

import (
	"fmt"

	e "github.com/abdullahnettoor/pass-gen/app/models/errors"
	"github.com/abdullahnettoor/pass-gen/app/models/req"
	"github.com/abdullahnettoor/pass-gen/app/models/res"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitRepository(DBs *gorm.DB) {
	db = DBs
}

const (
	userTable          = "pass_gen_users"
	passwordStoreTable = "pass_gen_password_stores"
)

func Signup(user *req.User) (*res.User, error) {
	var UserRes res.User
	query := fmt.Sprintf(`INSERT INTO %[1]s (user_name, password) 
		SELECT $1, $2 
		WHERE NOT EXISTS (SELECT 1 FROM %[1]s WHERE user_name = $1) 
		RETURNING *`, userTable)

	result := db.Raw(query, user.UserName, user.Password).Scan(&UserRes)
	if result.Error != nil {
		return nil, e.ErrDb
	}

	if result.RowsAffected == 0 {
		return nil, e.ErrUserConflict
	}
	return &UserRes, nil
}

func Login(user *req.User) (*res.LoginResponse, error) {
	var res res.LoginResponse
	query := fmt.Sprintf(`SELECT * FROM %[1]s WHERE user_name=$1`, userTable)
	result := db.Raw(query, user.UserName).Scan(&res)
	if result.Error != nil {
		return nil, e.ErrDb
	}

	if result.RowsAffected == 0 {
		return nil, e.ErrUserNotFound
	}
	return &res, nil
}
