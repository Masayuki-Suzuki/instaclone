package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var connectionError error

type SiteData struct {
	Title       string
	Description string
}

const (
	DRIVER_NAME = "mysql"
	DB_NAME     = "test_db"
	// user:password@tcp(container-name:port)/dbname *mysql is a default DB
	DATA_SOURCE_NAME = "root:thisisrootpassword@tcp(mariadb:3306)/mysql"
)

func init() {
	fmt.Println("Initializing DB...")
	db, connectionError = sql.Open(DRIVER_NAME, DATA_SOURCE_NAME)
	if connectionError != nil {
		log.Fatal("ERROR: Couldn't connect to the database:" + DB_NAME)
	}
}

func getIndexPage(w http.ResponseWriter, r *http.Request) {
	// Get DB informations
	rows, err := db.Query("SELECT SUBSTRING_INDEX(USER(), '@', -1) AS ip, @@hostname as hostname, @@port as port, DATABASE() as current_dtabase;")
	
	// Error Handling
	if err != nil {
		log.Print("Error: Execute database query: ", err)
		return
	}
	
	// Create string buffer
	var buffer bytes.Buffer
	
	// Create string
	for rows.Next() {
		var ip, hostname, port, currentDatabase string
		err = rows.Scan(&ip, &hostname, &port, &currentDatabase)
		buffer.WriteString("IP:: " + ip + "\nHostName:: " + hostname + "\nPort:: " + port)
	}
	
	fmt.Fprint(w, buffer.String())
}

func testHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this is a test.")
}

func main() {
	// Setup root directory
	//fs := http.FileServer(http.Dir("public"))
	
	router := mux.NewRouter().StrictSlash(true)

	// Add Routing: Access "/", shows public directory contents.
	router.HandleFunc("/", getIndexPage)
	router.HandleFunc("/test", testHandler)
	
	http.Handle("/", router)

	defer db.Close()
	// Shows message.
	fmt.Println("Server: Listening...")
	
	log.Fatal(http.ListenAndServe(":3000", nil))
}
