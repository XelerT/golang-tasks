//go:build !solution

package hotelbusiness

import (
	"sort"
)

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

func reserve_table2night(day int, guest_out_date int, calendar *map[int]int) {
	if day == guest_out_date {
		_, exists := (*calendar)[day]
		if !exists {
			(*calendar)[day] = 0
		}
	} else {
		(*calendar)[day]++
	}
}

func create_calendar(guests []Guest) map[int]int {
	calendar := map[int]int{}
	for i := range guests {
		for day := guests[i].CheckInDate; day <= guests[i].CheckOutDate; day++ {
			reserve_table2night(day, guests[i].CheckOutDate, &calendar)
		}
	}

	return calendar
}

func create_load(calendar map[int]int) []Load {
	load := []Load{}
	for day := range calendar {
		if calendar[day] == calendar[day-1] {
			continue
		}
		load = append(load, Load{StartDate: day, GuestCount: calendar[day]})
	}

	sort.SliceStable(load, func(p, q int) bool {
		return load[p].StartDate < load[q].StartDate
	})

	return load
}

func ComputeLoad(guests []Guest) []Load {
	return create_load(create_calendar(guests))
}
