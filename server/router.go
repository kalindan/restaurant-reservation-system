package server

import "github.com/go-chi/chi/v5"

func newRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/customer", func(r chi.Router) {
		r.Post("/register", register)
		r.Post("/", login)
		r.Get("/", logout)
	})
	r.Route("/reservation", func(r chi.Router) {
		r.Get("/make/", makeReservation)
		// r.Get("/", getReservations)
		// r.Delete("/", cancelReservation)
	})
	return r
}
