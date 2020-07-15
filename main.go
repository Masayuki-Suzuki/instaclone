package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
	
	firebase "firebase.google.com/go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
)

var db *sql.DB
var connectionError error

type SiteData struct {
	Title       	string
	Description 	string
	PageTitle		string
	ReactFilePath	string
	Users			[]User
}

type User struct {
	Username	string `json:"username"`
	Uid			string `json:"uid"`
	Email		string `json:"email"`
	FullName	string `json:"fullName"`
}

type ErrorMessage struct {
	ErrorMessage	string	`json:"errorMessage"`
}

const (
	DRIVER_NAME = "mysql"
	DB_NAME     = "instaclone"
	// user:password@tcp(container-name:port)/dbname *mysql is a default DB
	DATA_SOURCE_NAME = "root:thisisrootpassword@tcp(mariadb:3306)/" + DB_NAME
)

func init() {
	fmt.Println("Initializing DB...")
	db, connectionError = sql.Open(DRIVER_NAME, DATA_SOURCE_NAME)
	if connectionError != nil {
		log.Fatal("ERROR: Couldn't connect to the database:" + DB_NAME)
	}
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

func checkUserExistence(uid string) bool {
	var dbUids []string
	rows, err := db.Query("SELECT uid FROM users WHERE uid='" + uid + "'")
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
			dbUids = append(dbUids, dbUid)
		}
	} else {
		log.Println("rows are nil.")
		return false
	}
	
	return len(dbUids) > 0
}

func signUp(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var errorMessage string
	var user User
	
	decoder := json.NewDecoder(r.Body)
	
	err := decoder.Decode(&user)
	if err != nil {
		errorMessage = fmt.Sprintf("Error decoding json: %v", err)
		log.Println(errorMessage)
	}
	
	// If user hasn't existed.
	if !checkUserExistence(user.Uid) {
		opt := option.WithCredentialsFile("credentials/msa-instaclone-firebase-adminsdk-3gz0u-dfe554b4ec.json")
		app, err := firebase.NewApp(ctx, nil, opt)
		
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		
		// Setup firebase client
		client, err := app.Auth(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errorMessage = fmt.Sprintf("error getting Auth client: %v", err)
			log.Println(errorMessage)
		} else {
			// Get user data
			u, err := client.GetUser(ctx, user.Uid)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				errorMessage = fmt.Sprintf("error getting user %s: %v\n", user.Uid, err)
				log.Println(errorMessage)
			} else {
				var username string
				if user.Username != "" {
					username = user.Username
				} else {
					username = u.Email
				}
				// Send user data to database
				query := fmt.Sprintf("INSERT INTO users (uid, name, username, email, created_at) values('%v', '%v', '%v', '%v', '%v')",u.UID, u.DisplayName, username, u.Email, time.Now())
				_, err := db.Query(query)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					errorMessage = fmt.Sprintf("error create user data: %v\n", err)
					log.Printf(errorMessage)
				} else {
					returnData := User{
						Uid: u.UID,
						Email: u.Email,
						Username: username,
						FullName: u.DisplayName,
					}
					log.Println(json.NewEncoder(w).Encode(returnData))
				}
			}
		}
	} else {
		w.WriteHeader(http.StatusNotAcceptable)
		errorMessage = "User has already existed."
		log.Println(errorMessage)
	}
	
	if len(errorMessage) > 0 {
		log.Println(json.NewEncoder(w).Encode(ErrorMessage{ErrorMessage: errorMessage}))
	}
	log.Println(r.Body.Close())
}

// Shows user data from database. ** For testing purpose.
func getIndexPage(w http.ResponseWriter, r *http.Request) {
	// Get DB informations
	rows, err := db.Query("select uid, name, username, email from users")
	
	// Error Handling
	if err != nil {
		log.Print("Error: Execute database query: ", err)
		return
	}
	
	var users []User
	
	for rows.Next() {
		var uid, name, email, username  string
		err := rows.Scan(&uid, &name, &username, &email)
		
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(uid, name, username, email)
		users = append(users, User{ username, uid, email, name })
	}
	
	PageData := SiteData{
		Title:       "Test",
		Description: "This is test for DB connection, read and write.",
		PageTitle:   "DB Data",
		ReactFilePath: "",
		Users:       users,
	}
	
	t, err := template.ParseFiles("template/index.html.tpl")
	if err != nil {
		log.Printf("Could not find template file: %v", err)
	}
	
	if err := t.Execute(w, PageData); err != nil {
		log.Printf("Failed to execute template: %v", err)
	}
}

func paramTest (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["param"])
}

func main() {
	defer db.Close()
	
	// SetUp CORS
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:3000", "http://localhost:8088"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Authorization", "Origin", "Content-Type", "Accept"})
	
	router := mux.NewRouter().StrictSlash(true)
	
	// SignUp Api
	router.HandleFunc("/accounts/signup", authMiddleware(signUp)).Methods("POST")

	// Add Routing: Access "/", shows public directory contents.
	router.HandleFunc("/", getIndexPage)
	
	// Query Sample
	router.HandleFunc("/param-test/{param}", paramTest)
	
	// Shows message.
	fmt.Println("Server: Listening..")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)))
}
