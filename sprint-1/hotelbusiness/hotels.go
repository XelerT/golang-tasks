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

func create_calendar(guests []Guest) map[int] int {
	calendar := map[int] int {}
	for i := range guests {
		for day := guests[i].CheckInDate; day <= guests[i].CheckOutDate; day++ {
			if (day == guests[i].CheckOutDate) {
				calendar[day] = calendar[day]
			} else {
				calendar[day]++
			} 
		}
	}

	return calendar
}

func create_load(calendar map[int] int) []Load {
	load := []Load{}
	for day := range calendar {
		if calendar[day] == calendar[day - 1] {
			continue
		}
		load = append(load, Load{StartDate: day, GuestCount: calendar[day]})
	}

	sort.SliceStable(load,  func(p, q int) bool {  
		return load[p].StartDate < load[q].StartDate })

	return load
}

func ComputeLoad(guests []Guest) []Load {
	return create_load(create_calendar(guests))
}
