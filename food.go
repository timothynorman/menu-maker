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
		return food, fmt.Errorf("loadOneFood %v: %v", name, err)
	}

	return food, nil
}
