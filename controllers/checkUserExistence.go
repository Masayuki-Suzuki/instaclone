package controllers

import (
	"database/sql"
	"log"
	
	"github.com/Masayuki-Suzuki/instaclone/model"
)

func ScanUserRows(r *sql.Rows) bool {
	var data []string
	
	if r != nil {
		for r.Next() {
			var d string
			err := r.Scan(&d)
			if err != nil {
				log.Println("Scan error.")
			} else {
				data = append(data, d)
			}
		}
	} else {
		return false
	}
	
	return len(data) > 0
}

func CheckUserExistenceByUId(uid string) (bool, error) {
	// Get user from DB
	rows, err := model.GetUserByUId(uid)
	if err != nil {
		log.Println("DB Query Error: on checkUserExistence")
		log.Println(err)
	}
	
	return ScanUserRows(rows), err
}

func CheckUserExistenceByUserName(username string) (bool, error) {
	rows, err := model.GetUserByUserName(username)
	if err != nil {
		log.Println("DB Query Error: on checkUserExistence")
		log.Println(err)
	}
	
	return ScanUserRows(rows), err
}
