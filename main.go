package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB
var daysOfWeek = []string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

const (
	WEBSERVERPORT string = ":8000"
)

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
	fmt.Println("Starting Go webserver on port", WEBSERVERPORT)

	defaultHandler := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	}

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/menu", renderMenu)
	http.HandleFunc("/onemeal", updateOneMeal)

	log.Fatal(http.ListenAndServe(WEBSERVERPORT, nil))
}

func renderMenu(w http.ResponseWriter, r *http.Request) {
	menu := createMenu(7)

	data := struct {
		Menu [][]FoodItem
		Days []string
	}{
		Menu: menu,
		Days: daysOfWeek,
	}

	tmplFile := "menu.tmpl"
	t, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, data)
}

// updateOneMeal contains logic for endpoint to replace a single meal when it's re-roll button is clicked.
func updateOneMeal(w http.ResponseWriter, r *http.Request) {
	// get index of meal to be replaced from the menu.
	// sent by the button click.
	mealIndexStr := r.URL.Query().Get("mealIndex")
	mealIndex, err := strconv.Atoi(mealIndexStr)
	if err != nil || mealIndex < 0 {
		http.Error(w, "Invalid meal index", http.StatusBadRequest)
		return
	}

	// meal.tmpl is slighting different than menu.tmpl - it does not include range logic.
	tmplFile := "meal.tmpl"
	t, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// need to send the index and the new meal
	data := struct {
		Index int
		Meal  []FoodItem
		Day   string
	}{
		Index: mealIndex,
		Meal:  makeOneMeal(),
		Day:   daysOfWeek[mealIndex],
	}

	t.Execute(w, data)
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
