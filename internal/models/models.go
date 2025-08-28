package models

import "time"

type Account struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Is_active bool `json:"is_active"`
}

type AccountData struct{
	ID int
	Name string
	Amount float32
}

type TotalBalance struct{
	Balance float32
}

type AccountTransaction struct{
	ID int
	Date time.Time
	Category string
	Amount float32
}

type Category struct{
	ID int `json:"id"`
	Name string `json:"name"`
	type_of_category, is_public, Is_active bool
}