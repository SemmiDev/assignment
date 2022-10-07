package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	jumlahTicket := VIP + regular + student
	totalHargaTicket := (VIP * 30) + (regular * 20) + (student * 10)
	var totalBayar float32

	if totalHargaTicket >= 100 {
		// hari ganjil
		if day%2 != 0 {
			if jumlahTicket < 5 {
				totalBayar = float32(totalHargaTicket) - (float32(totalHargaTicket) * 0.15)
			} else {
				totalBayar = float32(totalHargaTicket) - (float32(totalHargaTicket) * 0.25)
			}
		}

		// hari genap
		if day%2 == 0 {
			if jumlahTicket < 5 {
				totalBayar = float32(totalHargaTicket) - (float32(totalHargaTicket) * 0.1)
			} else {
				totalBayar = float32(totalHargaTicket) - (float32(totalHargaTicket) * 0.2)
			}
		}
	} else {
		totalBayar = float32(totalHargaTicket)
	}

	return totalBayar
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(1, 1, 1, 20))
}
