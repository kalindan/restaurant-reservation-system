package server

import (
	"net/http"
	reservio "restaurant-project/reservation-system"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type server struct {
	rs reservio.ReservationSystem
}

func Serve() {
	r := chi.NewRouter()
	strg, _ := reservio.NewSqliteStorage()
	rs := reservio.NewReservationSystem(strg)
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(":3000", r)
}
