package config

import (
	"context"
	"log"
	
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var app *firebase.App

func FirebaseInit() {
	// Get Credentials File
	var opt = option.WithCredentialsFile("credentials/msa-instaclone-firebase-adminsdk-3gz0u-dfe554b4ec.json")
	// Get Firebase App
	a, err := firebase.NewApp(context.Background(), nil, opt)
	
	if err != nil {
		log.Printf("Error: Initializing Firebase App: %v\n", err)
	} else {
		app = a
	}
}

func GetFirebaseApp() *firebase.App {
	return app
}

func GetFirebaseAuthClient() (*auth.Client, error) {
	return app.Auth(context.Background())
}
