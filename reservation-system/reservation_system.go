package reservio

import (
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type reservationSystem struct {
	loggedCustomerName string
	dbh                dbHandler
}

func NewReservationSystem(dbh dbHandler) *reservationSystem {
	rs := &reservationSystem{
		dbh: dbh,
	}
	log.Print("New reservation system created")
	return rs
}

func (rs *reservationSystem) Register(name string, password string) error {
	if name == "" || password == "" {
		return errors.New("invalid name or password")
	}
	cs, _ := rs.dbh.getCustomer(name)
	if cs != nil {
		return errors.New("customer already registered")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	cs = newCustomer(name, string(hashed))
	err = rs.dbh.createCustomer(cs)
	if err != nil {
		return err
	}
	log.Printf("Customer %v registered with password %v", cs.name, password)
	return nil
}

func (rs *reservationSystem) Login(name string, password string) error {
	if rs.loggedCustomerName != "" {
		return errors.New("customer already logged in")
	}
	cs, err := rs.dbh.getCustomer(name)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(cs.password), []byte(password))
	if err != nil {
		return errors.New("invalid password")
	}
	rs.loggedCustomerName = cs.name
	log.Printf("Customer %v logged in", rs.loggedCustomerName)
	return nil
}

func (rs *reservationSystem) Logout() error {
	log.Printf("Customer %v logged out", rs.loggedCustomerName)
	rs.loggedCustomerName = ""
	return nil
}

func (rs *reservationSystem) MakeReservation(day int, hour int, duration int, persons int) error {
	if rs.loggedCustomerName == "" {
		return errors.New("not logged in")
	}
	if day > numOfDays {
		return errors.New("invalid day for reservation")
	}
	if hour < openingHour || hour > closingHour {
		return errors.New("invalid hour for reservation")
	}
	if hour+duration > closingHour {
		return errors.New("invalid duration for reservation")
	}
	if persons > maxPersons {
		return errors.New("invalid amount of persons for reservation")
	}
	tDay := day - 1
	tHour := hour - openingHour
	cs, err := rs.dbh.getCustomer(rs.loggedCustomerName)
	if err != nil {
		return err
	}
	tables, _ := rs.dbh.getTables() // place for possible cache
	for _, table := range *tables {
		if table.seats >= persons {
			isFree := true
			for i := tHour; i < tHour+duration; i++ {
				if table.days[tDay][i] {
					isFree = false
					break
				}
			}
			if isFree {
				for i := tHour; i < tHour+duration; i++ {
					table.days[tDay][i] = true
				}

				nr := newReservation(day, hour, duration, persons, cs.id, table.id)
				rs.dbh.updateTable(&table)
				rs.dbh.createReservation(nr)
				log.Printf("Reservation made at table %v in day %v", table.id, day)
				return nil
			}
		}
	}
	return errors.New("no available timeslot")
}

func (rs *reservationSystem) GetReservations() (*[]byte, error) {
	if rs.loggedCustomerName == "" {
		return nil, errors.New("not logged in")
	}
	cres, _ := rs.dbh.getCustomerReservations(rs.loggedCustomerName)
	var msg []byte
	msg = []byte(fmt.Sprintf("Reservations for customer %v\n", rs.loggedCustomerName))
	for i, res := range *cres {
		msg = append(msg, []byte(fmt.Sprintf("Reservation %v: Day %v, Hour %v, Duration %v, Persons %v, Table %v\n",
			i+1, res.day, res.hour, res.duration, res.persons, res.tableId))...)
	}
	return &msg, nil
}

func (rs *reservationSystem) CancelReservation(day int, tableId int) error {
	if rs.loggedCustomerName == "" {
		return errors.New("not logged in")
	}
	res, _ := rs.dbh.getCustomerReservation(rs.loggedCustomerName, day, tableId)
	err := rs.dbh.deleteReservation(rs.loggedCustomerName, day, tableId)
	if err != nil {
		return err
	}
	table, _ := rs.dbh.getTable(res.tableId)
	tHour := res.hour - openingHour
	tDay := res.day - 1
	for i := tHour; i < tHour+res.duration; i++ {
		table.days[tDay][i] = false
	}
	rs.dbh.updateTable(table)
	log.Printf("Reservation made at table %v in day %v was cancelled", tableId, day)
	return nil
}
