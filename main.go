package main

import (
	"log"
	"restaurant-project/reservation-system"
)

func main() {
	//restaurant : = reservio.NewRestaurant()
	dbh := reservio.NewTempDbHandle(10, 10)
	rs := reservio.NewReservationSystem(dbh)
	rs.Register("Milos", "12345678")
	cs, _ := rs.Dbh.GetCustomer("Milos")
	log.Print(*cs)

}
