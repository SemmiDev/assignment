package main

import "fmt"

func BMICalculator(gender string, height int) float64 {
	result := float64(height - 100)

	if gender == "laki-laki" {
		result = result - (result * 0.1)
	} else {
		result = result - (result * 0.15)
	}

	return result
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BMICalculator("laki-laki", 165))
	fmt.Println(BMICalculator("perempuan", 165))
}
