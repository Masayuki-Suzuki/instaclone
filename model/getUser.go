package model

import (
	"database/sql"
)

func GetUser(uid string) (*sql.Rows, error) {
	return db.Query("SELECT uid FROM users WHERE uid='" + uid + "'")
}

func GetAllUsers() (*sql.Rows, error) {
	return db.Query("select uid, name, username, email from users")
}
