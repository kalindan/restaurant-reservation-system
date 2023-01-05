package reservio

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const (
	createCustomersTable string = `CREATE TABLE IF NOT EXISTS customers (
		id TEXT PRIMARY KEY,
		name	TEXT UNIQUE,
		password	TEXT);`
	createReservationsTable string = `CREATE TABLE IF NOT EXISTS reservations (
		"id"		TEXT UNIQUE,
		"customerId"TEXT,
		"tableId"	INTEGER,
		"day"		INTEGER,
		"hour"		INTEGER,
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
		"day" 		INTEGER,
		"slot1"		INTEGER,
		"slot2"		INTEGER,
		"slot3"		INTEGER,
		"slot4"		INTEGER,
		"slot5"		INTEGER,
		"slot6"		INTEGER,
		"slot7"		INTEGER,
		"slot8"		INTEGER,
		"slot9"		INTEGER,
		"slot10"	INTEGER,
		"slot11"	INTEGER,
		"slot12"	INTEGER,
		FOREIGN KEY("tableId") REFERENCES "Table"("id")
	);`
	dbName string = "rs.db"
)

type sqliteStorage struct {
	db *sql.DB
}

func NewSqliteStorage() (*sqliteStorage, error) {
	var initRstTables bool
	if _, err := os.Stat(dbName); err != nil {
		initRstTables = true
	}
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}
	crArr := [4]string{createCustomersTable, createReservationsTable,
		createTablesTable, createTimeslotsTable}
	for _, create := range crArr {
		if _, err := db.Exec(create); err != nil {
			return nil, err
		}
	}
	if initRstTables {
		tables := newTables()
		for _, table := range tables {
			_, err := db.Exec("INSERT INTO tables VALUES(?,?);", table.id, table.seats)
			if err != nil {
				return nil, err
			}
			for _, tslts := range table.days {
				_, err := db.Exec("INSERT INTO timeslots VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?);", table.id, tslts.day,
					false, false, false, false, false, false, false, false, false, false, false, false)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	sqhl := &sqliteStorage{
		db: db,
	}
	return sqhl, nil
}

func (sq *sqliteStorage) createCustomer(cs *customer) error {
	_, err := sq.db.Exec("INSERT INTO customers VALUES(?,?,?);", cs.id, cs.name, cs.password)
	if err != nil {
		return err
	}
	return nil
}

func (sq *sqliteStorage) getCustomer(name string) (*customer, error) {
	row := sq.db.QueryRow("SELECT * FROM customers WHERE name=?", name)
	cs := &customer{}
	if err := row.Scan(&cs.id, &cs.name, &cs.password); err == sql.ErrNoRows {
		return nil, errors.New("customer not found")
	}
	return cs, nil
}

func (sq *sqliteStorage) getTable(id int) (*table, error) {
	tbRow := sq.db.QueryRow("SELECT id, seats FROM tables WHERE id=?", id)
	tb := table{}
	if err := tbRow.Scan(&tb.id, &tb.seats); err == sql.ErrNoRows {
		return nil, errors.New("table not found")
	}
	tslRows, err := sq.db.Query("SELECT * FROM timeslots WHERE tableId=?", id)
	if err != nil {
		return nil, err
	}
	defer tslRows.Close()
	for tslRows.Next() {
		tsl := timeslots{
			slots: make([]bool, closingHour-openingHour),
		}
		var tableId int
		err = tslRows.Scan(&tableId, &tsl.day, &tsl.slots[0], &tsl.slots[1],
			&tsl.slots[2], &tsl.slots[3], &tsl.slots[4], &tsl.slots[5], &tsl.slots[6],
			&tsl.slots[7], &tsl.slots[8], &tsl.slots[9], &tsl.slots[10], &tsl.slots[11])
		if err != nil {
			return nil, err
		}
		tb.days = append(tb.days, tsl)
	}

	return &tb, nil
}

func (sq *sqliteStorage) getTables() (*[]table, error) {
	tbRows, err := sq.db.Query("SELECT * FROM tables")
	if err != nil {
		return nil, err
	}
	defer tbRows.Close()
	tables := []table{}
	for tbRows.Next() {
		tbl := table{}
		err = tbRows.Scan(&tbl.id, &tbl.seats)
		if err != nil {
			return nil, err
		}
		tables = append(tables, tbl)
	}
	tslRows, err := sq.db.Query("SELECT * FROM timeslots")
	if err != nil {
		return nil, err
	}
	defer tslRows.Close()
	for tslRows.Next() {
		tsl := timeslots{
			slots: make([]bool, closingHour-openingHour),
		}
		var tableId int
		err = tslRows.Scan(&tableId, &tsl.day, &tsl.slots[0], &tsl.slots[1],
			&tsl.slots[2], &tsl.slots[3], &tsl.slots[4], &tsl.slots[5], &tsl.slots[6],
			&tsl.slots[7], &tsl.slots[8], &tsl.slots[9], &tsl.slots[10], &tsl.slots[11])
		if err != nil {
			return nil, err
		}
		for i, tb := range tables {
			if tableId == tb.id {
				tables[i].days = append(tables[i].days, tsl)
			}
		}

	}
	return &tables, nil
}

func (sq *sqliteStorage) updateTable(table *table) error {
	for _, tslts := range table.days {
		_, err := sq.db.Exec(`UPDATE timeslots SET slot1=?, slot2=?, slot3=?, 
			slot4=?, slot5=?, slot6=?, slot7=?, slot8=?, slot9=?, slot10=?, 
			slot11=?, slot12=? WHERE tableId=? AND day=?`, tslts.slots[0], tslts.slots[1],
			tslts.slots[2], tslts.slots[3], tslts.slots[4], tslts.slots[5], tslts.slots[6],
			tslts.slots[7], tslts.slots[8], tslts.slots[9], tslts.slots[10], tslts.slots[11],
			table.id, tslts.day)
		if err != nil {
			return err
		}
	}
	return nil
}

func (sq *sqliteStorage) createReservation(rs *reservation) error {
	_, err := sq.db.Exec("INSERT INTO reservations VALUES(?,?,?,?,?,?,?);", rs.id, rs.customerId,
		rs.tableId, rs.day, rs.hour, rs.duration, rs.persons)
	if err != nil {
		return err
	}
	return nil
}

func (sq *sqliteStorage) getCustomerReservation(name string, day int, tableId int) (*reservation, error) {
	cst, err := sq.getCustomer(name)
	if err != nil {
		return nil, err
	}
	rsrRow := sq.db.QueryRow("SELECT * FROM reservations WHERE customerId=? AND day=? and tableId=?", cst.id, day, tableId)
	rsr := &reservation{}
	if err := rsrRow.Scan(&rsr.id, &rsr.customerId, &rsr.tableId, &rsr.day, &rsr.hour, &rsr.duration,
		&rsr.persons); err == sql.ErrNoRows {
		return nil, errors.New("reservation not found")
	}
	return rsr, nil
}

func (sq *sqliteStorage) getCustomerReservations(name string) (*[]reservation, error) {
	cst, err := sq.getCustomer(name)
	if err != nil {
		return nil, err
	}
	rsrs := make([]reservation, 0)
	rsrRows, err := sq.db.Query("SELECT * FROM reservations WHERE customerId=?", cst.id)
	if err != nil {
		return nil, errors.New("no reservations available")

	}
	for rsrRows.Next() {
		rsr := reservation{}
		err = rsrRows.Scan(&rsr.id, &rsr.customerId, &rsr.tableId, &rsr.day,
			&rsr.hour, &rsr.duration, &rsr.persons)
		if err != nil {
			return nil, err
		}
		rsrs = append(rsrs, rsr)
	}
	return &rsrs, nil
}

func (sq *sqliteStorage) deleteReservation(name string, day int, tableId int) error {
	cst, err := sq.getCustomer(name)
	if err != nil {
		return err
	}
	_, err = sq.db.Exec("DELETE FROM reservations WHERE customerId=? AND day=? and tableId=?", cst.id, day, tableId)
	if err != nil {
		return err
	}
	return nil
}
