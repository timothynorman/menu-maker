package main

import (
	"database/sql"
	"fmt"
)

type FoodItem struct {
	Id          int
	Name        string
	Type        string
	LastUsed    sql.NullString
	Ingredients sql.NullString
	Bbq         bool
	Theme       string
}

func loadOneFoodByName(name string) (FoodItem, error) {
	var food FoodItem

	row := db.QueryRow("SELECT * FROM food WHERE name=?", name)
	if err := row.Scan(
		&food.Id,
		&food.Name,
		&food.Type,
		&food.LastUsed,
		&food.Ingredients,
		&food.Bbq,
		&food.Theme,
	); err != nil {
		if err == sql.ErrNoRows {
			return food, fmt.Errorf("no such food %v", name)
		}
		return food, fmt.Errorf("loadOneFoodByName %v: %v", name, err)
	}

	return food, nil
}

func loadOneFoodById(id int) (FoodItem, error) {
	var food FoodItem

	row := db.QueryRow("SELECT * FROM food WHERE id=?", id)
	if err := row.Scan(
		&food.Id,
		&food.Name,
		&food.Type,
		&food.LastUsed,
		&food.Ingredients,
		&food.Bbq,
		&food.Theme,
	); err != nil {
		if err == sql.ErrNoRows {
			return food, fmt.Errorf("no such food %v", id)
		}
		return food, fmt.Errorf("loadOneFoodById %v: %v", id, err)
	}

	return food, nil
}

func makeOneMeal() []FoodItem {
	var (
		completeMeal []FoodItem
		main         FoodItem
		side         FoodItem
		starch       FoodItem
	)

	main, _ = getMealOrMain()
	fmt.Printf("Added %v", main.Name)
	completeMeal = append(completeMeal, main)

	if main.Type == "main" {
		side, _ = getSideOrVeg()
		fmt.Printf(", and %v", side.Name)
		completeMeal = append(completeMeal, side)

		if side.Type == "veg" {
			starch, _ = getStarch()
			fmt.Printf(", and %v", starch.Name)
			completeMeal = append(completeMeal, starch)
		}
	}

	fmt.Printf(" to the meal.\n")

	return completeMeal
}

// createMenu makes unique meals up to the number of 'days'
func createMenu(days int) [][]FoodItem {
	var menu [][]FoodItem

	for i := 0; i < days; i++ {

		continueMeal := true
		for continueMeal {

			nextMeal := makeOneMeal()

			if !isRepeatedMeal(menu, nextMeal) {
				menu = append(menu, nextMeal)
				continueMeal = false
			}

		}

	}
	return menu
}

// isRepeatedMeal chekcs if the meal already exists in the menu.
func isRepeatedMeal(menu [][]FoodItem, newMeal []FoodItem) bool {
	for _, meal := range menu {
		if newMeal[0].Id == meal[0].Id {
			return true
		}
	}
	return false
}

// getMealOrMain returns a single random FoodItem that is either a 'Meal' or 'Main' type.
func getMealOrMain() (FoodItem, error) {
	row := db.QueryRow("SELECT * FROM food WHERE type='meal' OR type='main' ORDER BY RAND() LIMIT 1")
	return scanQuery(row)
}

// getSideOrBeg returns a single random FoodItem that is either a 'Side' or 'Veg' type.
func getSideOrVeg() (FoodItem, error) {
	row := db.QueryRow("SELECT * FROM food WHERE type='side' OR type='veg' ORDER BY RAND() LIMIT 1")
	return scanQuery(row)
}

// getStarch returns a single random FoodItem that is a 'Starch' type.
func getStarch() (FoodItem, error) {
	row := db.QueryRow("SELECT * FROM food WHERE type='starch' ORDER BY RAND() LIMIT 1")
	return scanQuery(row)
}

// scanQuery scans a row of FoodItem and assigns values to the struct.
func scanQuery(row *sql.Row) (FoodItem, error) {
	var food FoodItem

	if err := row.Scan(
		&food.Id,
		&food.Name,
		&food.Type,
		&food.LastUsed,
		&food.Ingredients,
		&food.Bbq,
		&food.Theme,
	); err != nil {
		if err == sql.ErrNoRows {
			return food, fmt.Errorf("no such food")
		}
		return food, fmt.Errorf("scanQuery: %v", err)
	}

	return food, nil
}
