package models

import "time"

type Account struct {
	ID int
	Name string
	Is_active bool
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
	ID int
	Name string
	type_of_category, is_public, Is_active bool
}