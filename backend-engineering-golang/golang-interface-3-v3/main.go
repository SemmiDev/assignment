package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	Hour   int
	Minute int
}

func mustValidNumber(str string) bool {
	for _, v := range str {
		if v < 48 || v > 57 {
			return false
		}
	}
	return true
}

func standartTime(time string) string {
	split := strings.Split(time, ":")
	hour, _ := strconv.Atoi(split[0])
	minute, _ := strconv.Atoi(split[1])

	if hour < 0 || hour > 23 || minute < 0 || minute > 59 {
		return "Invalid input"
	}

	if hour == 0 {
		return fmt.Sprintf("12:%02d AM", minute)
	}
	if hour == 12 {
		return fmt.Sprintf("12:%02d PM", minute)
	}
	if hour < 12 {
		return fmt.Sprintf("%02d:%02d AM", hour, minute)
	}
	return fmt.Sprintf("%02d:%02d PM", hour-12, minute)
}

func ChangeToStandartTime(time interface{}) string {
	switch time.(type) {
	case string:
		split := strings.Split(time.(string), ":")
		if len(split) != 2 {
			return "Invalid input"
		}
		if split[0] == "" || split[1] == "" {
			return "Invalid input"
		}
		if !(mustValidNumber(split[0]) && mustValidNumber(split[1])) {
			return "Invalid input"
		}
		return standartTime(time.(string))
	case []int:
		if len(time.([]int)) != 2 {
			return "Invalid input"
		}
		if time.([]int)[0] < 0 || time.([]int)[0] > 23 || time.([]int)[1] < 0 || time.([]int)[1] > 59 {
			return "Invalid input"
		}
		return standartTime(fmt.Sprintf("%02d:%02d", time.([]int)[0], time.([]int)[1]))
	case map[string]int:
		if len(time.(map[string]int)) != 2 {
			return "Invalid input"
		}
		if time.(map[string]int)["hour"] < 0 || time.(map[string]int)["hour"] > 23 || time.(map[string]int)["minute"] < 0 || time.(map[string]int)["minute"] > 59 {
			return "Invalid input"
		}
		// check if key must just contain hour or minute
		hour, ok := time.(map[string]int)["hour"]
		if !ok {
			return "Invalid input"
		}
		minute, ok := time.(map[string]int)["minute"]
		if !ok {
			return "Invalid input"
		}
		return standartTime(fmt.Sprintf("%02d:%02d", hour, minute))
	case Time:
		if time.(Time).Hour < 0 || time.(Time).Hour > 23 || time.(Time).Minute < 0 || time.(Time).Minute > 59 {
			return "Invalid input"
		}
		return standartTime(fmt.Sprintf("%02d:%02d", time.(Time).Hour, time.(Time).Minute))
	}
	return "Invalid input"
}

func main() {
	fmt.Println(ChangeToStandartTime("16:12"))
	fmt.Println(ChangeToStandartTime("13:00"))

	// fmt.Println(ChangeToStandartTime("16:00"))
	// fmt.Println(ChangeToStandartTime([]int{16, 0}))
	// fmt.Println(ChangeToStandartTime(map[string]int{"hour": 16, "minute": 0}))
	// fmt.Println(ChangeToStandartTime(Time{16, 0}))

}
