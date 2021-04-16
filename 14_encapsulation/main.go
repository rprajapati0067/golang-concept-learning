package main

import (
	"fmt"
	"log"

	"github.com/rprajapati0067/golang-concept-learning/14_encapsulation/calendar"
	"github.com/rprajapati0067/golang-concept-learning/14_encapsulation/event"
)

func main() {

	event := event.Event{}
	event.SetMonth(12)
	if err := event.SetTitle("Marriage"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(event)
	fmt.Println(event.Title())

	date := calendar.Date{}
	if err := date.SetYear(2021); err != nil {
		log.Fatal(err)
	}

	if err := date.SetMonth(04); err != nil {
		log.Fatal(err)
	}
	if err := date.SetDay(13); err != nil {
		log.Fatal(err)
	}

	//	fmt.Println((date))
	//	fmt.Printf("%d %d %d ", date.Year(), date.Month(), date.Day())
}
