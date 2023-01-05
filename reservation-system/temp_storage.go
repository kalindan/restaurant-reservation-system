package reservio

import (
	"errors"
)

type tempStorage struct {
	db *tempDb
}

func NewTempStorage(csAmount int, rsvAmount int) *tempStorage {
	tdbh := &tempStorage{
		db: newTempDb(csAmount, rsvAmount),
	}
	return tdbh
}

func (tdb *tempStorage) createCustomer(cs *customer) error {
	if _, err := tdb.getCustomer(cs.name); err == nil {
		return errors.New("customer already exists")
	}
	tdb.db.customers = append(tdb.db.customers, *cs)
	return nil
}

func (tdb *tempStorage) getCustomer(name string) (*customer, error) {
	for _, cs := range tdb.db.customers {
		if cs.name == name {
			return &cs, nil
		}
	}
	return nil, errors.New("customer not found")
}

func (tdb *tempStorage) getTable(id int) (*table, error) {
	for _, table := range tdb.db.tables {
		if table.id == id {
			return &table, nil
		}
	}
	return nil, errors.New("table not found")
}

func (tdb *tempStorage) getTables() (*[]table, error) {
	return &tdb.db.tables, nil
}

func (tdb *tempStorage) updateTable(table *table) error {
	for _, tb := range tdb.db.tables {
		if table.id == tb.id {
			tb = *table
		}
	}
	return nil
}

func (tdb *tempStorage) createReservation(rs *reservation) error {
	tdb.db.reservations = append(tdb.db.reservations, *rs)
	return nil
}

func (tdb *tempStorage) getCustomerReservation(name string, day int, tableId int) (*reservation, error) {
	cs, _ := tdb.getCustomer(name)
	for _, rs := range tdb.db.reservations {
		if rs.customerId == cs.id && rs.day == day && rs.tableId == tableId {
			return &rs, nil
		}
	}
	return nil, errors.New("reservation not found")
}

func (tdb *tempStorage) getCustomerReservations(name string) (*[]reservation, error) {
	var rss []reservation
	cs, _ := tdb.getCustomer(name)
	for _, rs := range tdb.db.reservations {
		if rs.customerId == cs.id {
			rss = append(rss, rs)
		}
	}
	return &rss, nil
}

func (tdb *tempStorage) deleteReservation(name string, day int, tableId int) error {
	cs, _ := tdb.getCustomer(name)
	for i, rs := range tdb.db.reservations {
		if rs.customerId == cs.id && rs.day == day && rs.tableId == tableId {
			tdb.db.reservations = append(tdb.db.reservations[:i], tdb.db.reservations[i+1:]...)
			return nil
		}
	}
	return errors.New("reservation not found")
}
