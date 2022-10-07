package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	// jumlahin total tiketnya
	jumlahTicket := VIP + regular + student

	// kaliin dengan harga masing-masing tiket, kemudian jumlahkan
	totalHargaTicket := (VIP * 30) + (regular * 20) + (student * 10)

	// siapin variable total bayar
	var totalBayar float32

	// cek jika total harga nya lebih besar atau sama dengan 100
	if totalHargaTicket >= 100 {
		// cek untuk hari ganjil
		if day%2 != 0 {
			// cek jumlah tiket
			if jumlahTicket < 5 {
				totalBayar = float32(totalHargaTicket) - (float32(totalHargaTicket) * 0.15) // 15%
			} else {
				totalBayar = float32(totalHargaTicket) - (float32(totalHargaTicket) * 0.25) // 25%
			}
		}

		// cek untuk hari genap
		if day%2 == 0 {
			// cek jumlah tiket
			if jumlahTicket < 5 {
				totalBayar = float32(totalHargaTicket) - (float32(totalHargaTicket) * 0.1) // 10%
			} else {
				totalBayar = float32(totalHargaTicket) - (float32(totalHargaTicket) * 0.2) // 20%
			}
		}
	} else {
		// kalau < 100 ga dpet diskon
		totalBayar = float32(totalHargaTicket)
	}

	return totalBayar
}

func main() {
	fmt.Println(GetTicketPrice(1, 1, 1, 20))
}
