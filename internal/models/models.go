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
	Type_of_category, Is_public, Is_active bool
}

type Transaction struct {
	ID int          
	Date time.Time  
	AccountID int   
	CategoryID int  
	Amount float64 
	Comment string   
}