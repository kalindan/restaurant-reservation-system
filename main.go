package main

import "restaurant-project/server"

func main() {
	// ui.Gui()
	server := server.NewReservationServer()
	server.Start(5000)
}
