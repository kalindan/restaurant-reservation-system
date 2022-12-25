package reservio

import (
	"golang.org/x/crypto/bcrypt"
)

type Reservation struct {
	id         int
	customerId int
	tableId    int
	day        int
	hour       int
	duration   int
	persons    int
}

type ReservationSystem struct {
	Dbh DbHandler
}

func NewReservationSystem(dbh DbHandler) *ReservationSystem {
	rs := &ReservationSystem{
		Dbh: dbh,
	}
	return rs
}

func (rs *ReservationSystem) Register(name string, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	cs := NewCustomer(name, string(hashed))
	err = rs.Dbh.CreateCustomer(cs)
	if err != nil {
		return err
	}
	return nil
}

func (rs *ReservationSystem) Login(name string, password string) error {
	//rs.Dbh.GetCustomer(name)
	return nil
}

// func (rs *ReservationSystem) MakeReservation(date int, hour int, duration float64, persons int) error {

// }

// func (rs *ReservationSystem) CancelReservation(id int) error {

// }
