package main

import (
	"fmt"
	"strings"
)

/*
	- split by space / comma, / semicolon (hint: strings.Contains, strings.Split)
	- tempShortestName := namesSplitted[0] // anggaplah index ke-0 adalah nama terpendek
	- bkin slice baru untuk menampung nama terpendek, masukin tempShortestName ke slice tsb
	- next, looping dari index ke-1 sampai akhir
	- jika panjang karakter dari index ke-i lebih kecil dari panjang karakter dari tempShortestName, maka tempShortestName diupdate dengan index ke-i
	- dan, slice yang beridis nama terpendek di assign/update lagi dengan tempShortestName
	- tapi kalau misalkan panjang karakter dari index ke-i sama dengan panjang karakter dari tempShortestName, maka nama dari index ke-i dimasukin ke slice nama terpendek (hint: append)
	- kenapa kita lakuin itu? karena bisa jadi ada lebih dari 1 nama terpendek yg panjangnya sama
	- terakhir, sort slice nama terpendek (hint: sort biasa pke loop / sort.Slice / sort.SliceStable, etc..)
*/

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
			fmt.Println(shortestNames)
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
	fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca ab Hamdan")) // "Tio"
	// FindShortestName("A,B,C,D,E")             // "Tia"
	// FindShortestName("Ari,Aru,Ara,Andi,Asik") // "Tia"
}
