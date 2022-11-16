package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	jumlahTicket := VIP + regular + student
	totalHargaTicket := float32((VIP * 30) + (regular * 20) + (student * 10))

	totalBayar := totalHargaTicket // default value if totalHargaTicket < 100

	if totalHargaTicket >= 100 {
		switch {
		case day%2 != 0:
			if jumlahTicket < 5 {
				totalBayar = totalHargaTicket - (totalHargaTicket * 0.15) // 15%
			} else {
				totalBayar = totalHargaTicket - (totalHargaTicket * 0.25) // 25%
			}
		case day%2 == 0:
			if jumlahTicket < 5 {
				totalBayar = totalHargaTicket - (totalHargaTicket * 0.1) // 10%
			} else {
				totalBayar = totalHargaTicket - (totalHargaTicket * 0.2) // 20%
			}
		}
	}
	return totalBayar
}

func main() {
	fmt.Println(GetTicketPrice(1, 1, 1, 20))
}
