package reservio

type Table struct {
	id        int
	seats     int
	timeslots []bool
}

func NewTables() []Table {
	tables := make([]Table, tableForTwo+tableForFour+tableForSix)
	i := 0
	for ; i < tableForTwo; i++ {
		tables[i].id = i
		tables[i].seats = 2
		tables[i].timeslots = make([]bool, openingHours)
	}
	for ; i < tableForTwo+tableForFour; i++ {
		tables[i].id = i
		tables[i].seats = 4
		tables[i].timeslots = make([]bool, openingHours)
	}
	for ; i < tableForTwo+tableForFour+tableForSix; i++ {
		tables[i].id = i
		tables[i].seats = 6
		tables[i].timeslots = make([]bool, openingHours)
	}
	return tables
}
