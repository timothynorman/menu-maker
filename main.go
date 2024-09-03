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

	h1 := func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, fmt.Sprint(makeOneMeal()))

		// this creates a tempalte based on what's written in 'index.html'
		tmpl := template.Must(template.ParseFiles("index.html"))

		Foods := makeOneMeal()

		TestValues := map[string][]FoodItem{
			"Foods": Foods,
		}
		tmpl.Execute(w, TestValues)
	}
	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8000", nil))

	// mealTest := makeOneMeal()
	// fmt.Println(mealTest)

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
