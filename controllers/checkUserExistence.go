package controllers

import (
	"log"
	
	"github.com/Masayuki-Suzuki/instaclone/model"
)

func CheckUserExistence(uid string) bool {
	var uids []string
	
	// Get user from DB
	rows, err := model.GetUser(uid)
	if err != nil {
		log.Println("DB Query Error: on checkUserExistence")
		log.Println(err)
	}
	
	if rows != nil {
		for rows.Next() {
			var dbUid string
			err := rows.Scan(&dbUid)
			if err != nil {
				log.Println("Scan error.")
				log.Println(err)
			}
			uids = append(uids, dbUid)
		}
	} else {
		return false
	}
	
	return len(uids) > 0
}
