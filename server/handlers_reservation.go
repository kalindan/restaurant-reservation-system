package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func makeReservation(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	cookie, err := r.Cookie("customer_name")
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(err.Error())
		return
	}
	dayStr := r.URL.Query().Get("day")
	day, _ := strconv.Atoi(dayStr)
	hourStr := r.URL.Query().Get("hour")
	hour, _ := strconv.Atoi(hourStr)
	log.Print(hour)
	durationStr := r.URL.Query().Get("duration")
	duration, _ := strconv.Atoi(durationStr)
	personsStr := r.URL.Query().Get("persons")
	persons, _ := strconv.Atoi(personsStr)
	err = rs.MakeReservation(cookie.Value, day, hour, duration, persons)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}
