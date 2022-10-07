package main

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
	return totalPerMonth(transpose((profitPerBranch)))
}
