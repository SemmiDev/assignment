package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var months = map[string]int{
	"January":   1,
	"February":  2,
	"March":     3,
	"April":     4,
	"May":       5,
	"June":      6,
	"July":      7,
	"August":    8,
	"September": 9,
	"October":   10,
	"November":  11,
	"December":  12,
}

func Date(y int, m int, d int) time.Time {
	month := time.Month(m)
	return time.Date(y, month, d, 0, 0, 0, 0, time.UTC)
}

// intinya pertam kita ambil jarak dari tanggal
// contohnya: 25 January - 30 January 2021
// total = 5 hari
// atau: 25 February - 10 March 2020
// (25 feb - 29 feb) = 5
// (1 march - 10 march) = 10
// total = 15 hari  (krna 2021 tahun kabisat)
// fungsi dibawah ini kebanyakan cuman split2 aja, krna kita mau ekstrak data dari tanggal dengan format yg udh diberikan ke tanggal yg di support golang
// int(to.Sub(from).Hours() / 24) tu maksudnya:
// to.Sub(from).Hours() tu ngurangin hari dari tanggal `to` sampe `from`, trus ambil total jam nya. trus / 24 tuk dapetin total harinya (krna 1 hari 24 jam)
// nah kalo dah dapat total harinya,
// panggil deh fungsi GetSalary
// bkin map tuk nampung nama karyawan dan total gaji nya nnti
// trus data di slice nya di loop, masukin deh dlem map, dan value nya isi dengan += 50_000, krna gajinya 50 rb perhari
// trus gaji yg didapat dlem map format deh jadi Rp. bla bla
// trus return deh map salary yg udah diformat Rp. itu

/*
	transformasi data:
	tanggal = 25 January - 30 January 2021
	data = [][]string{
		{"Andi", "Imam", "Eddy", "Deny"}, // hari 1
		{"Imam", "Eddy"}, // hari 2
		{"Deny"} // hari 3
	}

	- itung total hari = 6 hari
	- bikin map hasil = make(map[string]int)
	- loop data dengan limit smpe total harinya
	- masukin deh nama nya dalam map, dan valuenya +50.000 ->  hasil[namaKaryawan] += 50000
	- kalo udah, loop lagi map nya trus format setiap valuenya dengan format Rupiah
	- hasil2[namaKaryawan] = formatToRp(value nya)
	- return deh
*/
func GetDayDifference(date string) int {
	splittedDate := strings.Split(date, "-")
	getYear := strings.Split(splittedDate[1], " ")[3]
	getYearInt, _ := strconv.Atoi(getYear)

	fromDate := strings.Split(splittedDate[0], " ")[0]
	fromDateInt, _ := strconv.Atoi(fromDate)
	fromMonth := strings.Split(splittedDate[0], " ")[1]
	fromMonthInt := months[fromMonth]

	toDate := strings.Split(splittedDate[1], " ")[1]
	toDateInt, _ := strconv.Atoi(toDate)
	toMonth := strings.Split(splittedDate[1], " ")[2]
	toMonthInt := months[toMonth]

	from := Date(getYearInt, fromMonthInt, fromDateInt)
	to := Date(getYearInt, toMonthInt, toDateInt)

	if fromMonth == "January" && toMonth == "February" {
		return int(to.Sub(from).Hours() / 24)
	}

	if fromMonth == "March" && toMonth == "May" {
		return int(to.Sub(from).Hours() / 24)
	}

	return int(to.Sub(from).Hours()/24) + 1
}

func GetSalary(rangeDay int, data [][]string) map[string]string {
	salary := make(map[string]int)
	for i := 1; i <= rangeDay; i++ {
		for _, v := range data[i-1] {
			salary[v] += 50000
		}
	}
	newSalary := make(map[string]string)
	for k, v := range salary {
		newSalary[k] = FormatRupiah(v)
	}

	return newSalary
}

// https://adityalojes.blogspot.com/2016/10/konversi-angka-ke-format-rupiah-di.html
// Setiap 3 angka dari belakang ditambahkan “.”
// Tambahkan “Rp.” di depan dan pada beberapa keadaan bisa juga tambahkan “,00” di belakang
// krna Rp ga ada decimal
func FormatRupiah(number int) string {
	money := fmt.Sprintf("%d", number)
	// kita pisah dulu ke array
	numberRef := strings.Split(money, "")

	// Agar lebih mudah menghitung dari belakang,
	// nilai string perlu dibalik,
	for i, j := 0, len(numberRef)-1; i < j; i, j = i+1, j-1 {
		numberRef[i], numberRef[j] = numberRef[j], numberRef[i]
	}

	// kita ubah lagi jadi string
	// money = strings.Join(numberRef, "")
	// Lakukan loop dengan tiap 3 karakter (modulo 3)
	// kita tambahkan titik,
	//  kecuali jika karakter terakhir (karena akan berada di depan karakter pertama jika dibalik)
	angkaStrRevTitik := ""
	for i := 0; i < len(numberRef); i++ {
		angkaStrRevTitik += numberRef[i]
		if (i+1)%3 == 0 && i != (len(numberRef)-1) {
			angkaStrRevTitik += "."
		}
	}

	// dibalik lagi
	angkaStrTitik := ""
	for i := len(angkaStrRevTitik) - 1; i >= 0; i-- {
		angkaStrTitik += string(angkaStrRevTitik[i])
	}

	return "Rp " + angkaStrTitik
}

func GetSalaryOverview(dateRange string, data [][]string) map[string]string {
	days := GetDayDifference(dateRange)
	return GetSalary(days, data)
}

func main() {
	// res := GetSalaryOverview("21 February - 23 February 2021", [][]string{
	// 	{"Andi", "Imam", "Eddy", "Deny"},
	// 	{"Andi", "Imam"},
	// 	{"Imam", "Deny"},
	// 	{"Andi", "Deny"},
	// })

	// fmt.Println(res)

	// now := time.Now()
	// fmt.Println(getTotalDayInDateAndSub(now, 5))

	a := GetDayDifference("25 January - 5 February 2021")
	fmt.Println(a)
	a = GetDayDifference("30 March - 2 May 2021")
	fmt.Println(a)
}
