package model



import (
	"database/sql"
	"errors"
	"fmt"
)

func GetUserByUId(uid string) (*sql.Rows, error) {
	return db.Query("SELECT uid FROM users WHERE uid='" + uid + "'")
}

func GetUserByUserName(userName string) (*sql.Rows, error) {
	return db.Query("select username from users where username='" + userName + "'")
}

func GetUserByEmail(addr string) (*sql.Rows, error) {
	return db.Query("select email from users where email='"+ addr +"'")
}

func GetUserBy(col string, condition string) (*sql.Rows, error) {
	switch col {
	case "uid":
		return GetUserByUId(condition)
	case "username":
		return GetUserByUserName(condition)
	default:
		return nil, errors.New(fmt.Sprintf("Column name \"%v\" is invalid. Must choose from \"uid\" or \"username\".", col))
	}
}

func GetAllUsers() (*sql.Rows, error) {
	return db.Query("select uid, name, username, email, photo_url from users")
}
