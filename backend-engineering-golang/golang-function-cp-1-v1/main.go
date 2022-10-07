package main

import (
	"fmt"
)

func DateFormat(day, month, year int) string {
	months := map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	if day < 10 {
		return fmt.Sprintf("0%d-%s-%d", day, months[month], year)
	} else {
		return fmt.Sprintf("%d-%s-%d", day, months[month], year)
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(1, 1, 2012))
}
