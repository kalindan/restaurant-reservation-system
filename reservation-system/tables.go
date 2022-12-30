package reservio

type timeslots struct {
	day   int
	slots []bool
}
type table struct {
	id    int
	seats int
	days  []timeslots
}

func newTables() []table {
	tables := make([]table, tableForTwo+tableForFour+tableForSix)
	i := 0
	for ; i < tableForTwo; i++ {
		tables[i].id = i + 1
		tables[i].seats = 2
		tables[i].days = make([]timeslots, numOfDays)
		for j := range tables[i].days {
			tables[i].days[j].day = j + 1
			tables[i].days[j].slots = make([]bool, closingHour-openingHour)
		}
	}
	for ; i < tableForTwo+tableForFour; i++ {
		tables[i].id = i + 1
		tables[i].seats = 4
		tables[i].days = make([]timeslots, numOfDays)
		for j := range tables[i].days {
			tables[i].days[j].day = j + 1
			tables[i].days[j].slots = make([]bool, closingHour-openingHour)
		}
	}
	for ; i < tableForTwo+tableForFour+tableForSix; i++ {
		tables[i].id = i + 1
		tables[i].seats = 6
		tables[i].days = make([]timeslots, numOfDays)
		for j := range tables[i].days {
			tables[i].days[j].day = j + 1
			tables[i].days[j].slots = make([]bool, closingHour-openingHour)
		}
	}
	return tables
}
