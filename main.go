package main

import (
	"fmt"
	"log"
	"restaurant-project/reservation-system"
)

func main() {
	//restaurant : = reservio.NewRestaurant()
	dbh := reservio.NewTempDbHandle(10, 10)
	rs := reservio.NewReservationSystem(dbh)
	rs.Register("Milos", "12345678")
	_ = rs.Login("Milos", "12345678")
	rs.MakeReservation(2, 12, 2, 6)
	rs.MakeReservation(2, 12, 2, 6)
	rs.MakeReservation(2, 12, 2, 6)
	rs.MakeReservation(2, 12, 2, 6)
	rs.MakeReservation(2, 12, 2, 6)
	log.Print(rs.MakeReservation(2, 12, 2, 6))
	rss, _ := rs.GetReservations()
	fmt.Print(string(*rss))
	rs.Logout()
	log.Print(rs.MakeReservation(2, 12, 2, 3))
}
