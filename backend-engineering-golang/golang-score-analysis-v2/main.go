package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

type School struct {
	Name, Address string
	Grades        []int
}

func (s *School) AddGrade(grades ...int) {
	for _, v := range grades {
		s.Grades = append(s.Grades, v)
	}
}

func minMax(grades []int) (int, int) {
	min := grades[0]
	max := grades[0]

	for _, v := range grades {
		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	return min, max
}

func total(grades []int) int {
	total := 0

	for _, v := range grades {
		total += v
	}

	return total
}

func Analysis(s School) (float64, int, int) {
	if len(s.Grades) == 0 {
		return 0.0, 0, 0
	}

	min, max := minMax(s.Grades)
	total := total(s.Grades)
	avg := float64(total) / float64(len(s.Grades))
	return avg, min, max
}

// gunakan untuk melakukan debugging
func main() {
	avg, min, max := Analysis(School{
		Name:    "Imam Assidiqi School",
		Address: "Jl. Imam Assidiqi",
		Grades:  []int{100, 90, 100, 90, 100, 90},
	})

	fmt.Println(avg, min, max)
}
