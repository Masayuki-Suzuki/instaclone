package database

import (
"database/sql"
"fmt"
"log"
	
	"github.com/Masayuki-Suzuki/instaclone/config"
)

var Db *sql.DB
var connectionError error

func Init() {
	// Get config data
	c := config.GetConfig()
	fmt.Println("Initializing DB...")
	
	dbSourceName := c.GetString("database.data_source_name") + c.GetString("database.db_name")
	
	Db, connectionError = sql.Open(c.GetString("database.driver_name"), dbSourceName)
	if connectionError != nil {
		log.Fatal("ERROR: Couldn't connect to the database:" + c.GetString("database.db_name"))
	}
}

func Close() {
	Db.Close()
}

func GetDB() *sql.DB {
	return Db
}
