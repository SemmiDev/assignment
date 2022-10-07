package main

func SchedulableDays(date1 []int, date2 []int) []int {
	slots := make([]int, 0, len(date1)+len(date2))
	for _, v := range date1 {
		for _, v2 := range date2 {
			if v == v2 {
				slots = append(slots, v)
			}
		}
	}
	return slots
}
