package reservio

type TempDb struct {
	customers    []Customer
	reservations []Reservation
	tables       []Table
}

func NewTempDb(csAmount int, rsvAmount int) *TempDb {
	tdb := &TempDb{
		customers:    make([]Customer, csAmount),
		reservations: make([]Reservation, rsvAmount),
		tables:       NewTables(),
	}
	return tdb
}
