package reservio

type tempDb struct {
	customers    []customer
	reservations []reservation
	tables       []table
}

func newTempDb(csAmount int, rsvAmount int) *tempDb {
	tdb := &tempDb{
		customers:    make([]customer, csAmount),
		reservations: make([]reservation, rsvAmount),
		tables:       newTables(),
	}
	return tdb
}
