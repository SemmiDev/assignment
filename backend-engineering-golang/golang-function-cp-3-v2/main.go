package main

import (
	"strings"
)

func FindShortestName(names string) string {
	namesSplitted := strings.Split(names, " ") // default split by space
	if strings.Contains(names, ";") {
		namesSplitted = strings.Split(names, ";") // split by semicolon
	}
	if strings.Contains(names, ",") {
		namesSplitted = strings.Split(names, ",") // split by comma
	}

	shortestName := namesSplitted[0]
	shortestNames := []string{shortestName}

	for i := 1; i < len(namesSplitted); i++ {
		if len(namesSplitted[i]) < len(shortestName) {
			shortestName = namesSplitted[i]
			shortestNames = []string{shortestName}
		} else if len(namesSplitted[i]) == len(shortestName) {
			shortestNames = append(shortestNames, namesSplitted[i])
		}
	}

	// sort shortestNames
	for i := 0; i < len(shortestNames); i++ {
		for j := i + 1; j < len(shortestNames); j++ {
			if shortestNames[i] > shortestNames[j] {
				shortestNames[i], shortestNames[j] = shortestNames[j], shortestNames[i]
			}
		}
	}

	return shortestNames[0]
}

// gunakan untuk melakukan debug
func main() {
	// fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan")) // "Tio"
	FindShortestName("A,B,C,D,E")             // "Tia"
	FindShortestName("Ari,Aru,Ara,Andi,Asik") // "Tia"
}
