package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

const ()

func main() {

	// Load environment variables
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal(envErr)
	}

	// Assign variables to connection configs
	cfg := &mysql.Config{
		User:   os.Getenv("DB_USERNAME"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "menu_maker",
	}

	ConnectToDatabase(cfg)

	// Start Go webserver
	fmt.Println("Starting Go webserver.")

	defaultHandler := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	}

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/menu", renderMenu)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func renderMenu(w http.ResponseWriter, r *http.Request) {
	menu := createMenu(7)

	tmplFile := "meal.tmpl"
	t, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, menu)
}

func ConnectToDatabase(cfg *mysql.Config) error {
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Printf("could not open connection to database: %v", err)
		return err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("could not ping database: %v", err)
		return err
	}

	fmt.Println("Connected to database")

	return nil
}
