package reservio

type tableStorage interface {
	getTable(id int) (*table, error)
	getTables() (*[]table, error)
	updateTable(table *table) error
}

type reservationStorage interface {
	createReservation(rs *reservation) error
	getCustomerReservation(name string, day int, tableId int) (*reservation, error)
	getCustomerReservations(name string) (*[]reservation, error)
	deleteReservation(name string, day int, tableId int) error
}

type customerStorage interface {
	createCustomer(cs *customer) error
	getCustomer(name string) (*customer, error)
}

type storage interface {
	tableStorage
	reservationStorage
	customerStorage
}
