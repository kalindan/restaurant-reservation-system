package reservio

import "github.com/google/uuid"

type Customer struct {
	id       uuid.UUID
	name     string
	password string
}

func NewCustomer(name string, password string) *Customer {
	cst := &Customer{
		id:       uuid.New(),
		name:     name,
		password: password}
	return cst
}
