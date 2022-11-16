package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PopulationData(data []string) (people []map[string]any) {
	if len(data) == 0 {
		return []map[string]any{}
	}

	for _, v := range data {
		dataSplit := strings.Split(v, ";")
		age, _ := strconv.Atoi(dataSplit[1])

		personMap := map[string]any{
			"name":    dataSplit[0],
			"age":     age,
			"address": dataSplit[2],
		}

		height := dataSplit[3]
		if height != "" {
			heightFloat, _ := strconv.ParseFloat(height, 64)
			personMap["height"] = heightFloat
		}

		isMarried := dataSplit[4]
		if isMarried != "" {
			isMarriedBool, _ := strconv.ParseBool(isMarried)
			personMap["isMarried"] = isMarriedBool
		}

		people = append(people, personMap)
	}
	return people
}

func main() {
	result := PopulationData([]string{
		"Budi;23;Jakarta;;",
		"Joko;30;Bandung;;true",
		"Susi;25;Bogor;165.42;",
	})

	fmt.Println(result)
}
