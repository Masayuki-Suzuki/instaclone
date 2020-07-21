package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	
	"github.com/Masayuki-Suzuki/instaclone/controllers"
	"github.com/Masayuki-Suzuki/instaclone/middleware"
	"github.com/Masayuki-Suzuki/instaclone/model"
	"github.com/Masayuki-Suzuki/instaclone/types"
	"github.com/gorilla/mux"
)

type SiteData struct {
	Title       	string
	Description 	string
	PageTitle		string
	ReactFilePath	string
	Users			[]types.User
}

// Shows user data from database. ** For testing purpose.
func getIndexPage(w http.ResponseWriter, r *http.Request) {
	// Get DB information
	rows, err := model.GetAllUsers()
	
	// Error Handling
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print("Error: Execute database query: ", err)
		return
	}
	
	var users []types.User
	
	for rows.Next() {
		var uid, name, email, username  string
		err := rows.Scan(&uid, &name, &username, &email)
	
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(uid, name, username, email)
		users = append(users, types.User{ username, uid, email, name })
	}
	
	PageData := SiteData{
		Title:       "Test",
		Description: "This is test for DB connection, read and write.",
		PageTitle:   "DB Data",
		ReactFilePath: "",
		Users:       users,
	}
	
	t, err := template.ParseFiles("templates/index.html.tpl")
	
	if err != nil {
		log.Printf("Could not find template file: %v", err)
	}
	
	if err := t.Execute(w, PageData); err != nil {
		log.Printf("Failed to execute template: %v", err)
	}
	
}

// Shows text from url query. ** For testing purpose.
func paramTest (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["param"])
}


func NewRouter() *mux.Router {
	// Router Setup
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/accounts/signup", middleware.Auth(controllers.SignUp)).Methods("POST")
	router.HandleFunc("/", getIndexPage)
	// Query Sample
	router.HandleFunc("/param-test/{param}", paramTest)
	
	return router
}
