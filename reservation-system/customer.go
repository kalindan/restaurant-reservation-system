package reservio

type Customer struct {
	id       int
	name     string
	password string
}

func NewCustomer(name string, password string) Customer {
	cst := Customer{
		name:     name,
		password: password}
	return cst
}
