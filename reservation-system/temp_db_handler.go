package reservio

import (
	"errors"
)

type tempDbHandler struct {
	db *tempDb
}

func NewTempDbHandler(csAmount int, rsvAmount int) *tempDbHandler {
	tdbh := &tempDbHandler{
		db: newTempDb(csAmount, rsvAmount),
	}
	return tdbh
}

func (tdb *tempDbHandler) createCustomer(cs *customer) error {
	if _, err := tdb.getCustomer(cs.name); err == nil {
		return errors.New("customer already exists")
	}
	tdb.db.customers = append(tdb.db.customers, *cs)
	return nil
}

func (tdb *tempDbHandler) getCustomer(name string) (*customer, error) {
	for _, cs := range tdb.db.customers {
		if cs.name == name {
			return &cs, nil
		}
	}
	return nil, errors.New("customer not found")
}

func (tdb *tempDbHandler) getTable(id int) (*table, error) {
	for _, table := range tdb.db.tables {
		if table.id == id {
			return &table, nil
		}
	}
	return nil, errors.New("table not found")
}

func (tdb *tempDbHandler) getTables() (*[]table, error) {
	return &tdb.db.tables, nil
}

func (tdb *tempDbHandler) updateTable(table *table) error {
	for _, tb := range tdb.db.tables {
		if table.id == tb.id {
			tb = *table
		}
	}
	return nil
}

func (tdb *tempDbHandler) createReservation(rs *reservation) error {
	tdb.db.reservations = append(tdb.db.reservations, *rs)
	return nil
}

func (tdb *tempDbHandler) getCustomerReservation(name string, day int, tableId int) (*reservation, error) {
	cs, _ := tdb.getCustomer(name)
	for _, rs := range tdb.db.reservations {
		if rs.customerId == cs.id && rs.day == day && rs.tableId == tableId {
			return &rs, nil
		}
	}
	return nil, errors.New("reservation not found")
}

func (tdb *tempDbHandler) getCustomerReservations(name string) (*[]reservation, error) {
	var rss []reservation
	cs, _ := tdb.getCustomer(name)
	for _, rs := range tdb.db.reservations {
		if rs.customerId == cs.id {
			rss = append(rss, rs)
		}
	}
	return &rss, nil
}

func (tdb *tempDbHandler) deleteReservation(name string, day int, tableId int) error {
	cs, _ := tdb.getCustomer(name)
	for i, rs := range tdb.db.reservations {
		if rs.customerId == cs.id && rs.day == day && rs.tableId == tableId {
			tdb.db.reservations = append(tdb.db.reservations[:i], tdb.db.reservations[i+1:]...)
			return nil
		}
	}
	return errors.New("reservation not found")
}
