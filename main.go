package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var db *sql.DB
var connectionError error

type User struct {
	Id			int
	First_name	string
	Last_name	string
	Gender		int
	Age			int
}

type SiteData struct {
	Title       	string
	Description 	string
	PageTitle		string
	ReactFilePath	string
	Users			[]User
}

const (
	DRIVER_NAME = "mysql"
	DB_NAME     = "test_db"
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

//func addDummyUser(w http.ResponseWriter, r *http.Request) {
//	_, err := db.Query("insert into users (first_name, last_name, gender, age) values('Robert', 'Baratheon', 0 , 40) ")
//	if err != nil {
//		log.Print("Error: Execute database query: ", err)
//		return
//	}
//	log.Fatalln(fmt.Fprint(w, "Added User? Should Back to index page."))
//}

func getIndexPage(w http.ResponseWriter, r *http.Request) {
	// Get DB informations
	rows, err := db.Query("select * from users")
	
	// Error Handling
	if err != nil {
		log.Print("Error: Execute database query: ", err)
		return
	}
	
	var users []User
	
	for rows.Next() {
		var first_name, last_name string
		var id, gender, age int
		err := rows.Scan(&id, &first_name, &last_name, &gender, &age)
		
		if err != nil {
			log.Fatalln(err)
		}
		users = append(users, User{ id, first_name, last_name, gender, age })
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
		log.Fatal("Could not find template file: %v", err)
	}
	
	if err := t.Execute(w, PageData); err != nil {
		log.Printf("Failed to execute template: %v", err)
	}
	
	// Create string buffer
	//var buffer bytes.Buffer
	
	// Create string
	//for rows.Next() {
	//	var ip, hostname, port, currentDatabase string
	//	err = rows.Scan(&ip, &hostname, &port, &currentDatabase)
	//	buffer.WriteString("IP:: " + ip + "\nHostName:: " + hostname + "\nPort:: " + port)
	//}
	
	//fmt.Fprint(w, "this is a test")
}

func testHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this is a test.")
}

func main() {
	defer db.Close()
	
	// SetUp CORS
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhots:8088"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Authorization"})
	
	
	// Setup root directory
	//fs := http.FileServer(http.Dir("public"))
	
	router := mux.NewRouter().StrictSlash(true)

	// Add Routing: Access "/", shows public directory contents.
	router.HandleFunc("/", getIndexPage)
	//router.HandleFunc("/add", addDummyUser)
	router.HandleFunc("/test", testHandler)
	
	// Shows message.
	fmt.Println("Server: Listening..")
	
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)))
}
