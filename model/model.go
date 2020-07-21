package model

import (
	"database/sql"
	
	"github.com/Masayuki-Suzuki/instaclone/database"
)

var db *sql.DB

func Init() {
	db = database.GetDB()
}
