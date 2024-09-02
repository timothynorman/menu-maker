package main

import (
	"database/sql"
	"fmt"
)

var menu [][]FoodItem

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
	var completeMeal []FoodItem

	main, _ := getMealOrMain()
	side, _ := getSideOrVeg()

	completeMeal = append(completeMeal, main, side)
	fmt.Printf("Added %v and %v to the meal.\n", main.Name, side.Name)

	if side.Type == "veg" {
		starch, _ := getStarch()
		completeMeal = append(completeMeal, starch)
		fmt.Printf("Added %v to the meal.\n", starch.Name)
	}

	return completeMeal
}

// getMealOrMain returns a single random FoodItem that is either a 'Meal' or 'Main' type.
func getMealOrMain() (FoodItem, error) {
	var food FoodItem

	row := db.QueryRow("SELECT * FROM food WHERE type='meal' OR type='main' ORDER BY RAND() LIMIT 1")
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
		return food, fmt.Errorf("getMealOrMain: %v", err)
	}

	return food, nil
}

// getSideOrBeg returns a single random FoodItem that is either a 'Side' or 'Veg' type.
func getSideOrVeg() (FoodItem, error) {
	var food FoodItem

	row := db.QueryRow("SELECT * FROM food WHERE type='side' OR type='veg' ORDER BY RAND() LIMIT 1")
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
		return food, fmt.Errorf("getSideOrVeg: %v", err)
	}

	return food, nil
}

// getStarch returns a single random FoodItem that is a 'Starch' type.
func getStarch() (FoodItem, error) {
	var food FoodItem

	row := db.QueryRow("SELECT * FROM food WHERE type='starch' ORDER BY RAND() LIMIT 1")
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
		return food, fmt.Errorf("getStarch: %v", err)
	}

	return food, nil
}
