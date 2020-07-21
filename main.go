package main

import (
	"github.com/Masayuki-Suzuki/instaclone/config"
	"github.com/Masayuki-Suzuki/instaclone/database"
	"github.com/Masayuki-Suzuki/instaclone/model"
	"github.com/Masayuki-Suzuki/instaclone/server"
	
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	defer database.Close()
	
	config.Init("development")
	config.FirebaseInit()
	database.Init()
	model.Init()
	
	server.Init()
}
