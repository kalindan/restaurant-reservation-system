package main

import (
	"log"
	reservio "restaurant-project/reservation-system"
	rsui "restaurant-project/ui"
)

func main() {
	_, err := reservio.NewSqliteHandler()
	if err != nil {
		log.Print(err.Error())
	}
	rsui.Gui()
}
