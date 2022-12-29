package reservio

import "github.com/google/uuid"

type customer struct {
	id       uuid.UUID
	name     string
	password string
}

func newCustomer(name string, password string) *customer {
	cst := &customer{
		id:       uuid.New(),
		name:     name,
		password: password}
	return cst
}
