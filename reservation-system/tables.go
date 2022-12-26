package reservio

type timeslots []bool
type Table struct {
	id    int
	seats int
	days  []timeslots
}

func NewTables() []Table {
	tables := make([]Table, tableForTwo+tableForFour+tableForSix)
	i := 0
	for ; i < tableForTwo; i++ {
		tables[i].id = i
		tables[i].seats = 2
		tables[i].days = make([]timeslots, numOfDays)
		for j := range tables[i].days {
			tables[i].days[j] = make([]bool, closingHour-openingHour)
		}
	}
	for ; i < tableForTwo+tableForFour; i++ {
		tables[i].id = i
		tables[i].seats = 4
		tables[i].days = make([]timeslots, numOfDays)
		for j := range tables[i].days {
			tables[i].days[j] = make([]bool, closingHour-openingHour)
		}
	}
	for ; i < tableForTwo+tableForFour+tableForSix; i++ {
		tables[i].id = i
		tables[i].seats = 6
		tables[i].days = make([]timeslots, numOfDays)
		for j := range tables[i].days {
			tables[i].days[j] = make([]bool, closingHour-openingHour)
		}
	}
	return tables
}
