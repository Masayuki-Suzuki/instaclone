package model

import (
	"errors"
	"fmt"
	"time"
	
	"github.com/Masayuki-Suzuki/instaclone/database"
	"github.com/Masayuki-Suzuki/instaclone/types"
)

func CreateNewUser(u *types.User) error  {
	db := database.GetDB()
	query := fmt.Sprintf(
		"INSERT INTO users (uid, name, username, email, created_at, photo_url) " +
			"values('%v', '%v', '%v', '%v', '%v', '%v')",
		u.Uid, u.FullName, u.Username, u.Email, time.Now(), u.PhotoUrl)
	_, err := db.Query(query)
	
	if err != nil {
		return errors.New(fmt.Sprintf("error create user data: %v\n", err))
	} else {
		return nil
	}
}
