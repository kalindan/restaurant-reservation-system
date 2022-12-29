package reservio

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteHandler struct {
	db *sql.DB
}

func NewSqliteHandler() (*sqliteHandler, error) {
	db, err := sql.Open("sqlite3", "rs.db")
	if err != nil {
		return nil, err
	}
	sqhl := &sqliteHandler{
		db: db,
	}
	return sqhl, nil
}

func (sq *sqliteHandler) createCustomer(cs *customer) error {
	_, err := sq.db.Exec("INSERT INTO Customer VALUES(?,?,?);", cs.id, cs.name, cs.password)
	if err != nil {
		return err
	}
	return nil
}

func (sq *sqliteHandler) getCustomer(name string) (*customer, error) {
	row := sq.db.QueryRow("SELECT * FROM Customer WHERE name=?", name)
	cs := customer{}
	if err := row.Scan(&cs.id, &cs.name, &cs.password); err == sql.ErrNoRows {
		return nil, errors.New("customer not found")
	}
	return &cs, nil
}

func (sq *sqliteHandler) getTable(id int) (*table, error) {
	panic("not implemented") // TODO: Implement
}

func (sq *sqliteHandler) getTables() (*[]table, error) {
	panic("not implemented") // TODO: Implement
}

func (sq *sqliteHandler) updateTable(tables *table) error {
	panic("not implemented") // TODO: Implement
}

func (sq *sqliteHandler) createReservation(rs *reservation) error {
	panic("not implemented") // TODO: Implement
}

func (sq *sqliteHandler) getCustomerReservation(name string, day int, tableId int) (*reservation, error) {
	panic("not implemented") // TODO: Implement
}

func (sq *sqliteHandler) getCustomerReservations(name string) (*[]reservation, error) {
	panic("not implemented") // TODO: Implement
}

func (sq *sqliteHandler) deleteReservation(name string, day int, tableId int) error {
	panic("not implemented") // TODO: Implement
}
