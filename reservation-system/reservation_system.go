package reservio

import (
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type ReservationSystem struct {
	loggedCustomerName string
	dbh                storage
}

func NewReservationSystem(dbh storage) *ReservationSystem {
	rs := &ReservationSystem{
		dbh: dbh,
	}
	log.Print("New reservation system created")
	return rs
}

func (rs *ReservationSystem) Register(name string, password string) error {
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

func (rs *ReservationSystem) Login(name string, password string) error {
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

func (rs *ReservationSystem) Logout() error {
	log.Printf("Customer %v logged out", rs.loggedCustomerName)
	rs.loggedCustomerName = ""
	return nil
}

func (rs *ReservationSystem) MakeReservation(day int, hour int, duration int, persons int) error {
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
	tHour := hour - openingHour
	cs, err := rs.dbh.getCustomer(rs.loggedCustomerName)
	if err != nil {
		return err
	}
	tables, err := rs.dbh.getTables() // place for possible cache
	if err != nil {
		return err
	}
	for _, table := range *tables {
		if table.seats >= persons {
			isFree := true
			for _, timeslots := range table.days {
				for i := tHour; i < tHour+duration; i++ {
					if timeslots.day == day && timeslots.slots[i] {
						isFree = false
						break
					}
				}
			}
			if isFree {
				for j, timeslots := range table.days {
					for i := tHour; i < tHour+duration; i++ {
						if timeslots.day == day {
							table.days[j].slots[i] = true
						}
					}
				}
				nr := newReservation(day, hour, duration, persons, cs.id, table.id)
				err := rs.dbh.updateTable(&table)
				if err != nil {
					return err
				}
				err = rs.dbh.createReservation(nr)
				if err != nil {
					return err
				}
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
	cres, err := rs.dbh.getCustomerReservations(rs.loggedCustomerName)
	if err != nil {
		return nil, err
	}
	var msg []byte
	msg = []byte(fmt.Sprintf("Reservations for customer %v\n", rs.loggedCustomerName))
	for i, res := range *cres {
		msg = append(msg, []byte(fmt.Sprintf("Reservation %v: Day %v, Hour %v, Duration %v, Persons %v, Table %v\n",
			i+1, res.day, res.hour, res.duration, res.persons, res.tableId))...)
	}
	return &msg, nil
}

func (rs *ReservationSystem) CancelReservation(day int, tableId int) error {
	if rs.loggedCustomerName == "" {
		return errors.New("not logged in")
	}
	res, err := rs.dbh.getCustomerReservation(rs.loggedCustomerName, day, tableId)
	if err != nil {
		return err
	}
	err = rs.dbh.deleteReservation(rs.loggedCustomerName, day, tableId)
	if err != nil {
		return err
	}
	table, err := rs.dbh.getTable(res.tableId)
	if err != nil {
		return err
	}
	tHour := res.hour - openingHour
	for i := tHour; i < tHour+res.duration; i++ {
		for j, timeslots := range table.days {
			if timeslots.day == day {
				table.days[j].slots[i] = false
			}
		}
	}
	err = rs.dbh.updateTable(table)
	if err != nil {
		return err
	}
	log.Printf("Reservation made at table %v in day %v was cancelled", tableId, day)
	return nil
}
