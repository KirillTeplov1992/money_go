package main

import (
	"log"
	application "money/internal/app"
	"money/internal/app"
)

func main(){
	new_config := app.NewConfig()
	new_app := application.New(new_config)
	if err := new_app.Start(); err != nil {
		log.Fatal(err)
	} 

}