package main

import (
	"database/sql"
	"fmt"
	"log"
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

	// test, _ := loadOneFoodByName("spaghetti")
	// fmt.Println(test)

	// test2, _ := loadOneFoodById(7)
	// fmt.Println(test2)

	// fmt.Println(test2.Name)

	mealTest := makeOneMeal()
	fmt.Println(mealTest)

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
