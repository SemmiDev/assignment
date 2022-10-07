package main

import (
	"fmt"
	"sort"
)

func SchedulableDays(villager [][]int) []int {
	result := make([]int, 0, 100)
	maps := make(map[int]int)

	if len(villager) == 1 {
		return villager[0]
	}

	for _, v := range villager {
		for _, v2 := range v {
			maps[v2]++
		}
	}

	villagerLen := len(villager)
	for k, v := range maps {
		if v == villagerLen {
			result = append(result, k)
		}
	}

	sort.Ints(result)
	return result
}

func main() {
	data := [][]int{
		{10, 11, 12, 13, 14, 15, 16, 20, 21, 22, 23, 24, 25, 26, 27, 28},
		{10, 11, 12, 13, 14, 15, 16, 20, 21, 22, 23, 24, 25, 26, 27, 28},
		{10, 11, 12, 13, 14, 15, 16, 20, 21, 22, 23, 24, 25, 26, 27, 28},
		{10, 11, 12, 13, 14, 15, 16, 20, 21, 22, 23, 24, 25, 26, 27, 28},
		{10, 11, 12, 13, 14, 15, 16, 20, 21, 22, 23, 24, 25, 26, 27, 28},
	}

	fmt.Println(SchedulableDays(data))
}
