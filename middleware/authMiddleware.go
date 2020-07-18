package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type ErrorMessage struct {
	ErrorMessage	string	`json:"errorMessage"`
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Setup Firebase SDK
		opt := option.WithCredentialsFile("credentials/msa-instaclone-firebase-adminsdk-3gz0u-dfe554b4ec.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		
		if err != nil {
			fmt.Printf("error initializing app: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			// Initialise Auth Func.
			auth, err := app.Auth(context.Background())
			if err != nil {
				fmt.Printf("error initializing auth function: %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				// Get JWT from Client
				authHeader := r.Header.Get("Authorization")
				idToken := strings.Replace(authHeader, "Bearer ", "", 1)
				
				// Validate JWT
				_, err := auth.VerifyIDToken(context.Background(), idToken)
				if err != nil {
					// If JWT is invalid, work error handling instead of Handler.
					fmt.Printf("error verifying ID Token: %v\n", err)
					w.WriteHeader(http.StatusUnauthorized)
					
					errorMessage := ErrorMessage{ErrorMessage: "error verifying ID Token"}
					
					log.Println(json.NewEncoder(w).Encode(errorMessage))
					return
				}
				next.ServeHTTP(w, r)
			}
		}
	}
}
