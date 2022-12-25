package reservio

import (
	"errors"
)

type DbHandler interface {
	CreateCustomer(cs Customer) error
	GetCustomer(name string) (*Customer, error)
	MakeReservation()
	CancelReservation()
	SetupRestaurant()
}

type TempDbHandle struct {
	db *TempDb
}

func NewTempDbHandle(csAmount int, rsvAmount int) *TempDbHandle {
	tdbh := &TempDbHandle{
		db: NewTempDb(csAmount, rsvAmount),
	}
	return tdbh
}

func (tdb *TempDbHandle) CreateCustomer(cs Customer) error {
	for _, dcs := range tdb.db.customers {
		if dcs.name == cs.name {
			return errors.New("Customer already exists")
		}
	}
	tdb.db.customers = append(tdb.db.customers, cs)
	return nil
}

func (tdb *TempDbHandle) GetCustomer(name string) (*Customer, error) {
	for _, cs := range tdb.db.customers {
		if cs.name == name {
			return &cs, nil
		}
	}
	return nil, errors.New("Customer not found")
}

func (tdb *TempDbHandle) MakeReservation() {

}

func (tdb *TempDbHandle) CancelReservation() {
}

func (tdb *TempDbHandle) SetupRestaurant() {
}
