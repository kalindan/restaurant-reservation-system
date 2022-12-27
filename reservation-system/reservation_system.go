package reservio

import (
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type ReservationSystem struct {
	loggedCustomerName string
	dbh                DbHandler
}

func NewReservationSystem(dbh DbHandler) *ReservationSystem {
	rs := &ReservationSystem{
		dbh: dbh,
	}
	log.Print("New reservation system created")
	return rs
}

func (rs *ReservationSystem) Register(name string, password string) error {
	cs, _ := rs.dbh.GetCustomer(name)
	if cs != nil {
		return errors.New("Customer already registered")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	cs = NewCustomer(name, string(hashed))
	err = rs.dbh.CreateCustomer(cs)
	if err != nil {
		return err
	}
	log.Printf("Customer %v registered", cs.name)
	return nil
}

func (rs *ReservationSystem) Login(name string, password string) error {
	cs, err := rs.dbh.GetCustomer(name)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(cs.password), []byte(password))
	if err != nil {
		return err
	}
	rs.loggedCustomerName = cs.name
	log.Printf("Customer %v logged in", rs.loggedCustomerName)
	return nil
}

func (rs *ReservationSystem) Logout() error {
	log.Printf("Customer %v logged out", rs.loggedCustomerName)
	rs.loggedCustomerName = ""
	return nil
}

func (rs *ReservationSystem) MakeReservation(day int, hour int, duration int, persons int) error {
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
	if rs.loggedCustomerName == "" {
		return errors.New("not logged in")
	}
	tDay := day - 1
	tHour := hour - openingHour
	cs, err := rs.dbh.GetCustomer(rs.loggedCustomerName)
	if err != nil {
		return err
	}
	tables, _ := rs.dbh.GetTables() // place for possible cache
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

				nr := NewReservation(day, hour, duration, persons, cs.id, table.id)
				rs.dbh.UpdateTables(tables)
				rs.dbh.CreateReservation(nr)
				log.Printf("Reservation made at table %v in day %v", table.id, day)
				return nil
			}
		}
	}
	return errors.New("no available timeslot")
}

func (rs *ReservationSystem) GetReservations() (*[]byte, error) {
	if rs.loggedCustomerName == "" {
		return nil, errors.New("not logged in")
	}
	cres, _ := rs.dbh.GetCustomerReservations(rs.loggedCustomerName)
	var msg []byte
	msg = []byte(fmt.Sprintf("Reservations for customer %v\n", rs.loggedCustomerName))
	for _, res := range *cres {
		msg = append(msg, []byte(fmt.Sprintf("Day %v, Hour %v, Duration %v, Persons %v, Table %v\n",
			res.day, res.hour, res.duration, res.persons, res.tableId))...)
	}
	return &msg, nil
}

func (rs *ReservationSystem) CancelReservation(day int, tableId int) error {
	if rs.loggedCustomerName == "" {
		return errors.New("not logged in")
	}
	res, _ := rs.dbh.GetCustomerReservation(rs.loggedCustomerName, day, tableId)
	err := rs.dbh.DeleteReservation(rs.loggedCustomerName, day, tableId)
	if err != nil {
		return err
	}
	table, _ := rs.dbh.GetTable(res.tableId)
	tHour := res.hour - openingHour
	tDay := res.day - 1
	for i := tHour; i < tHour+res.duration; i++ {
		table.days[tDay][i] = false
	}
	rs.dbh.UpdateTable(table)
	log.Printf("Reservation made at table %v in day %v was cancelled", tableId, day)
	return nil
}
