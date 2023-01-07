package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type reservationServer struct {
	rtr *chi.Mux
}

func NewReservationServer() *reservationServer {
	sr := &reservationServer{
		rtr: chi.NewRouter(),
	}
	return sr
}

func (rsrv *reservationServer) Start(port int) {
	rsrv.rtr.Use(middleware.Logger)
	rsrv.rtr.Mount("/", newRouter())
	log.Printf("Reservation server running on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), rsrv.rtr))

}
