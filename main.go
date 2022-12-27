package main

import (
	"fmt"
	"log"
	reservio "restaurant-project/reservation-system"
	"time"
)

func main() {
	start := time.Now()
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
	rs.CancelReservation(2, 17)
	rss, _ = rs.GetReservations()
	fmt.Print(string(*rss))
	rs.MakeReservation(2, 12, 2, 6)
	rss, _ = rs.GetReservations()
	fmt.Print(string(*rss))
	rs.Logout()
	log.Print(rs.MakeReservation(2, 12, 2, 3))
	elapsed := time.Since(start)
	log.Printf("Operation took %s", elapsed)
}
