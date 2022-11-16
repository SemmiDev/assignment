package main

import "fmt"

/*
[80 120 180 220]
[60 100 120 130]
------------------+
[120 220 300 350]
*/

func transpose(data [][]int) [][]int {
	xl := len(data[0])
	yl := len(data)
	result := make([][]int, xl)
	for i := range result {
		result[i] = make([]int, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = data[j][i]
		}
	}
	fmt.Println(data)
	fmt.Println(result)
	return result
}

func totalPerMonth(data [][]int) []int {
	result := make([]int, 0, len(data))
	for _, v := range data {
		temp := 0
		for _, v2 := range v {
			temp += v2
		}
		result = append(result, temp)
	}
	return result
}

func CountProfit(data [][][2]int) []int {
	if len(data) == 0 {
		return []int{}
	}
	profitPerBranch := make([][]int, len(data))
	for i, branch := range data {
		for _, perMonth := range branch {
			profit := perMonth[0] - perMonth[1]
			profitPerBranch[i] = append(profitPerBranch[i], profit)
		}
	}

	/*
	data = [][][2]int{
		{
			{1000, 800},
			{700, 500},
			{100, 50}
		},
		{
			{1000, 800},
			{900, 200},
			{500, 200}
		},
		{
			{1000, 900},
			{900, 200},
			{500, 200}
		}

	- pertama kita hitung dulu profit tiap bulan di tiap cabang,
	  hasil perhitungan profitnya kita masukin aja dlam slice baru

	  jadinya = [
		[200 200 50] // profit pada bulan 1
		[200 700 300] // profit pada bulan 2
		[100 700 300] // profit pada bulan 3
	 ]

	- kita mau gabungin profit pada bulan sama pada setiap cabang
	  caranya ? transpose matrix
	  jadinya = [
		[200 200 100]
		[200 700 700]
		[50 300 300]
	  ]

	 - baru deh kita totalin tiap baris nya
		200 + 200 + 100 = 500
		200 + 700 + 700 = 1600
		50 + 300 + 300 = 650

	  - return [500,1600,650]
		*/









	return totalPerMonth(transpose((profitPerBranch)))
}



func main() {
	r := CountProfit([][][2]int{
		{{1000, 800}, {700, 500}, {100, 50}},
		{{1000, 800}, {900, 200}, {500, 200}},
		{{1000, 900}, {900, 200}, {500, 200}}})

		fmt.Println(r)
}

// []int{500, 350, 500, 50}
