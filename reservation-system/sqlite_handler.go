package reservio

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

const (
	createCustomersTable string = `CREATE TABLE IF NOT EXISTS customers (
		id TEXT PRIMARY KEY,
		name	TEXT UNIQUE,
		password	TEXT);`
	createReservationsTable string = `CREATE TABLE IF NOT EXISTS reservations (
		"id"	TEXT UNIQUE,
		"customerId"	TEXT,
		"tableId"	INTEGER,
		"day"	INTEGER,
		"hour"	INTEGER,
		"duration"	INTEGER,
		"persons"	INTEGER,
		FOREIGN KEY("tableId") REFERENCES "Table"("id"),
		FOREIGN KEY("customerId") REFERENCES "Customer"("id"),
		PRIMARY KEY("id"));`
	createTablesTable string = `CREATE TABLE IF NOT EXISTS tables (
		"id"	INTEGER UNIQUE,
		"seats"	INTEGER,
		PRIMARY KEY("id"));`
	createTimeslotsTable string = `CREATE TABLE IF NOT EXISTS timeslots (
		"tableId"	INTEGER,
		"1"	INTEGER,
		"2"	INTEGER,
		"3"	INTEGER,
		"4"	INTEGER,
		"5"	INTEGER,
		"6"	INTEGER,
		"7"	INTEGER,
		"8"	INTEGER,
		"9"	INTEGER,
		"10"	INTEGER,
		"11"	INTEGER,
		"12"	INTEGER,
		FOREIGN KEY("tableId") REFERENCES "Table"("id")
	);`
)

type sqliteHandler struct {
	db *sql.DB
}

func NewSqliteHandler() (*sqliteHandler, error) {
	db, err := sql.Open("sqlite3", "rs.db")
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(createCustomersTable); err != nil {
		return nil, err
	}
	if _, err := db.Exec(createReservationsTable); err != nil {
		return nil, err
	}
	if _, err := db.Exec(createTablesTable); err != nil {
		return nil, err
	}
	if _, err := db.Exec(createTimeslotsTable); err != nil {
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
