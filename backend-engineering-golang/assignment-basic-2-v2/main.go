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

func IsLeapYear(year string) bool {
	yearInt, _ := time.Parse("2006", year)
	if yearInt.Year()%4 == 0 {
		if yearInt.Year()%100 == 0 {
			return yearInt.Year()%400 == 0
		}
		return true
	}
	return false
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
