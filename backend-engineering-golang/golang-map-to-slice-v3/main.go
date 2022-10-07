package main

import "fmt"

func MapToSlice(mapData map[string]string) [][]string {
	if len(mapData) == 0 {
		return [][]string{}
	}

	transform := make([][]string, len(mapData))
	i := 0
	for key, value := range mapData {
		transform[i] = []string{key, value}
		i++
	}
	return transform
}

func main() {
	var matkul string
	var nilai int

	// 90-100 A
	// 80-90 B
	// 70-79 C
	// 60-69 D
	// < 60 E

	fmt.Print("Masukkan nama matkul: ")
	_, _ = fmt.Scan(&matkul)

	fmt.Print("Masukkan nilai: ")
	_, _ = fmt.Scan(&nilai)

	if nilai >= 90 && nilai <= 100 {
		fmt.Println("Nilai A")
	} else if nilai >= 80 && nilai <= 89 {
		fmt.Println("Nilai B")
	} else if nilai >= 70 && nilai <= 79 {
		fmt.Println("Nilai C")
	} else if nilai >= 60 && nilai <= 69 {
		fmt.Println("Nilai D")
	} else {
		fmt.Println("Nilai E")
	}
}
