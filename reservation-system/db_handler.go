package reservio

import (
	"errors"
)

type DbHandler interface {
	CreateCustomer(cs *Customer) error
	GetCustomer(name string) (*Customer, error)
	GetTable(id int) (*Table, error)
	GetTables() (*[]Table, error)
	UpdateTables(tables *[]Table) error
	CreateReservation(rs *Reservation) error
	GetCustomerReservations(name string) (*[]Reservation, error)
	DeleteReservation()
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

func (tdb *TempDbHandle) CreateCustomer(cs *Customer) error {
	if _, err := tdb.GetCustomer(cs.name); err == nil {
		return errors.New("Customer already exists")
	}
	tdb.db.customers = append(tdb.db.customers, *cs)
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

func (tdb *TempDbHandle) GetTable(id int) (*Table, error) {
	for _, table := range tdb.db.tables {
		if table.id == id {
			return &table, nil
		}
	}
	return nil, errors.New("Table not found")
}

func (tdb *TempDbHandle) GetTables() (*[]Table, error) {
	return &tdb.db.tables, nil
}

func (tdb *TempDbHandle) UpdateTables(tables *[]Table) error {
	tdb.db.tables = *tables
	return nil
}

func (tdb *TempDbHandle) CreateReservation(rs *Reservation) error {
	tdb.db.reservations = append(tdb.db.reservations, *rs)
	return nil
}

func (tdb *TempDbHandle) GetCustomerReservations(name string) (*[]Reservation, error) {
	var rss []Reservation
	cs, _ := tdb.GetCustomer(name)
	for _, rs := range tdb.db.reservations {
		if rs.customerId == cs.id {
			rss = append(rss, rs)
		}
	}
	return &rss, nil
}

func (tdb *TempDbHandle) DeleteReservation() {
}
