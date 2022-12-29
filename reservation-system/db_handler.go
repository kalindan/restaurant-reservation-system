package reservio

type dbHandler interface {
	createCustomer(cs *customer) error
	getCustomer(name string) (*customer, error)
	getTable(id int) (*table, error)
	getTables() (*[]table, error)
	updateTable(tables *table) error
	createReservation(rs *reservation) error
	getCustomerReservation(name string, day int, tableId int) (*reservation, error)
	getCustomerReservations(name string) (*[]reservation, error)
	deleteReservation(name string, day int, tableId int) error
}
