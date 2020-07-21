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
	
	// Json setup
	decoder := json.NewDecoder(r.Body)
	
	err := decoder.Decode(&formData)
	if err != nil {
		errorMessage = fmt.Sprintf("Error decoding json: %v", err)
		log.Println(errorMessage)
	}
	
	// Set Headers
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
	// Check username has already existed or not if user chose email sign up.
	if formData.EmailSignUp {
		userExistenceByUserName, err := CheckUserExistenceByUserName(formData.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(json.NewEncoder(w).Encode(
				types.ErrorMessage{ErrorMessage: "DB Query Error: on CheckUserExistenceByUserName."}))
		}
		// username has already existed.
		if userExistenceByUserName {
			w.WriteHeader(http.StatusConflict)
			log.Println(json.NewEncoder(w).Encode(
				types.ErrorMessage{ErrorMessage: "Cannot use the user name. It's already used."}))
		}
	}
	
	// Check uid has already existed or not.
	userExistenceByUId, err := CheckUserExistenceByUId(formData.Uid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(json.NewEncoder(w).Encode(
			types.ErrorMessage{ErrorMessage: "DB Query Error: on CheckUserExistenceByUId"}))
	}
	
	// If user hasn't existed.
	if !userExistenceByUId {
		// Setup firebase client
		client, err := config.GetFirebaseAuthClient()
		if err != nil {
			// Return 500 error
			w.WriteHeader(http.StatusInternalServerError)
			errorMessage = fmt.Sprintf("error getting Auth client: %v", err)
			log.Println(json.NewEncoder(w).Encode(types.ErrorMessage{ErrorMessage: errorMessage}))
		}
		
		// Get user data from firebase
		u, err := client.GetUser(ctx, formData.Uid)
		if err != nil {
			// Return 500 error
			w.WriteHeader(http.StatusInternalServerError)
			errorMessage = fmt.Sprintf("error getting user %s: %v\n", formData.Uid, err)
			log.Println(json.NewEncoder(w).Encode(types.ErrorMessage{ErrorMessage: errorMessage}))
		}
		
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
			PhotoUrl: formData.PhotoUrl,
		}
		
		// Create new user on DB
		err = model.CreateNewUser(&userData)
		if err != nil {
			// Return 500 error
			w.WriteHeader(http.StatusInternalServerError)
			errorMessage = fmt.Sprintf("error create user data: %v\n", err)
			log.Println(json.NewEncoder(w).Encode(types.ErrorMessage{ErrorMessage: errorMessage}))
		}
		
		// Completed correctly.
		w.WriteHeader(http.StatusOK)
		log.Println(json.NewEncoder(w).Encode(userData))
	} else {
		// uid has already existed.
		w.WriteHeader(http.StatusConflict)
		log.Println(json.NewEncoder(w).Encode(types.ErrorMessage{ErrorMessage: "User has already existed."}))
	}
	
	log.Println(r.Body.Close())
}
