package main

import (
	"log"
	"time"

	"site-accessibility/helpers"
)

func main() {
	go loopTheCheck()
	router := SetupRouter()
	router.Run()
}

func loopTheCheck() {
	for {
		helpers.CheckNewData()
		log.Println("Data updated!")
		time.Sleep(1 * time.Minute)
	}
}
