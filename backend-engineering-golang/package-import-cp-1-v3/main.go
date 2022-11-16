package main

import (
	"a21hc3NpZ25tZW50/internal"
	"fmt"
	"strconv"
	"strings"
)

func AdvanceCalculator(calculate string) float32 {
	if calculate == "" {
		return 0.0
	}
	splitExp := strings.Split(calculate, " ")

	num := splitExp[0]
	numFloat, _ := strconv.ParseFloat(num, 32)

	if len(splitExp) == 1 {
		return float32(numFloat)
	}

	/*
		3 * 4  / 2  + 10  - 5 = length 9
		  1 2 [3] 4 [5] 6 [7] 8

		  1 + 2 - 3 + 4 * 5
		  base = 1

		  loop dari angka index ke 1 sampai length - 1
		  ambil nih mula2 index ke 1
		  otomatis dia operator
		  ambil index ke 2
		  otomatis dia angka


	*/

	calc := internal.NewCalculator(float32(numFloat))
	for i := 1; i < len(splitExp)-1; i += 2 {
		switch splitExp[i] {
		case "+":
			num = splitExp[i+1]
			numFloat, _ = strconv.ParseFloat(num, 32)
			calc.Add(float32(numFloat))
		case "-":
			num = splitExp[i+1]
			numFloat, _ = strconv.ParseFloat(num, 32)
			calc.Subtract(float32(numFloat))
		case "*":
			num = splitExp[i+1]
			numFloat, _ = strconv.ParseFloat(num, 32)
			calc.Multiply(float32(numFloat))
		case "/":
			num = splitExp[i+1]
			numFloat, _ = strconv.ParseFloat(num, 32)
			calc.Divide(float32(numFloat))
		}
	}

	return calc.Result()
}

func main() {
	res := AdvanceCalculator("3 * 4 / 2 + 10 - 5")
	fmt.Println(res)
	res = AdvanceCalculator("10 / 4 + 100")
	fmt.Println(res)
	res = AdvanceCalculator("10 + 10 + 10 + 10 + 12 + 12 + 12 + 12")
	fmt.Println(res)
}
