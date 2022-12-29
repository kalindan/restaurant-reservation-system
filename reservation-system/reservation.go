package reservio

import "github.com/google/uuid"

type reservation struct {
	id         uuid.UUID
	customerId uuid.UUID
	tableId    int
	day        int
	hour       int
	duration   int
	persons    int
}

func newReservation(day int, hour int, duration int, persons int, customerId uuid.UUID, tableId int) *reservation {
	nr := &reservation{
		id:         uuid.New(),
		customerId: customerId,
		tableId:    tableId,
		day:        day,
		hour:       hour,
		duration:   duration,
		persons:    persons,
	}
	return nr
}
