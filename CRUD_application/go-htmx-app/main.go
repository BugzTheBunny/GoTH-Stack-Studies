package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func getConnectionString() string {
	userName := "root"
	password := "root"
	host := "127.0.0.1"
	port := "3333"
	databaseName := "testdb"
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", userName, password, host, port, databaseName)
}

var db *sql.DB
var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func initDB() {
	var err error
	driverName := "mysql"
	connectionString := getConnectionString()
	db, err = sql.Open(driverName, connectionString)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

}

func main() {
	initDB()
	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler)

	http.ListenAndServe(":3000", router)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "home.html", nil)

	if err != nil {
		http.Error(w, "Error executing template"+err.Error(), http.StatusInternalServerError)
	}
}
