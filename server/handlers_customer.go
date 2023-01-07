package server

import (
	"encoding/json"
	"log"
	"net/http"
	reservio "restaurant-project/reservation-system"
)

var dbh, _ = reservio.NewSqliteStorage()
var rs = reservio.NewReservationSystem(dbh)

func register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	customer := new(customer)
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := decoder.Decode(customer)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(err.Error())
		return
	}
	err = rs.Register(customer.Name, customer.Password)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	customer := new(customer)
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := decoder.Decode(customer)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(err.Error())
		return
	}
	if err = rs.Login(customer.Name, customer.Password); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(err.Error())
		return
	}
	cookie := http.Cookie{
		Name:     "customer_name",
		Value:    customer.Name,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}

func logout(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	cookie, err := r.Cookie("customer_name")
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(err.Error())
		return
	}
	if err = rs.Logout(cookie.Value); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(err.Error())
		return
	}

	emptyCookie := http.Cookie{
		Name:     "customer_name",
		Value:    "",
		HttpOnly: true,
	}
	http.SetCookie(w, &emptyCookie)
	w.WriteHeader(http.StatusOK)
}
