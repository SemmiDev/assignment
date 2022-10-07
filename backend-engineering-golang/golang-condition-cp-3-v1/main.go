package main

import "fmt"

func GetPredicate(math, science, english, indonesia int) string {
	avg := (math + science + english + indonesia) / 4
	switch {
	case avg == 100:
		return "Sempurna"
	case avg >= 90:
		return "Sangat Baik"
	case avg >= 80:
		return "Baik"
	case avg >= 70:
		return "Cukup"
	case avg >= 60:
		return "Kurang"
	case avg < 60:
		return "Sangat kurang"
	default:
		return "Nilai tidak valid"
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetPredicate(50, 80, 100, 60))
}
