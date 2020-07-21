package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	
	"github.com/Masayuki-Suzuki/instaclone/config"
	"github.com/Masayuki-Suzuki/instaclone/model"
	"github.com/Masayuki-Suzuki/instaclone/types"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var errorMessage string
	var formData types.SignUpForm
	
	decoder := json.NewDecoder(r.Body)
	
	err := decoder.Decode(&formData)
	if err != nil {
		errorMessage = fmt.Sprintf("Error decoding json: %v", err)
		log.Println(errorMessage)
	}
	
	// If user hasn't existed.
	if !CheckUserExistence(formData.Uid) {
		// Set Headers
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		
		// Setup firebase client
		client, err := config.GetFirebaseAuthClient()
		if err != nil {
			// Return 500 error
			w.WriteHeader(http.StatusInternalServerError)
			errorMessage = fmt.Sprintf("error getting Auth client: %v", err)
			log.Println(errorMessage)
		} else {
			// Get user data from firebase
			u, err := client.GetUser(ctx, formData.Uid)
			if err != nil {
				// Return 500 error
				w.WriteHeader(http.StatusInternalServerError)
				errorMessage = fmt.Sprintf("error getting user %s: %v\n", formData.Uid, err)
				log.Println(errorMessage)
			} else {
				
				var username, fullName, email string
				
				// If user sign up with email
				if formData.EmailSignUp {
					username = formData.Username
					fullName = formData.FullName
					email = formData.Email
				} else {
					// If user sign up with google
					username = u.Email
					fullName = u.DisplayName
					email = u.Email
				}
				
				// Create user data for inserting it to DB.
				userData := types.User{
					Uid: u.UID,
					Username: username,
					FullName: fullName,
					Email: email,
				}
				
				// Create new user on DB
				err := model.CreateNewUser(userData)
				if err != nil {
					// Return 500 error
					w.WriteHeader(http.StatusInternalServerError)
					errorMessage = fmt.Sprintf("error create user data: %v\n", err)
					log.Printf(errorMessage)
				} else {
					w.WriteHeader(http.StatusOK)
					log.Println(json.NewEncoder(w).Encode(userData))
				}
			}
		}
	} else {
		w.WriteHeader(http.StatusNotAcceptable)
		errorMessage = "User has already existed."
		log.Println(errorMessage)
	}
	
	if len(errorMessage) > 0 {
		log.Println(json.NewEncoder(w).Encode(types.ErrorMessage{ErrorMessage: errorMessage}))
	}
	log.Println(r.Body.Close())
}
