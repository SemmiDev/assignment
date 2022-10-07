package main

import (
	"fmt"
	"strconv"
	"strings"
)

var LocationCode = map[string]string{
	"JKT": "JKT",
	"BDG": "BDG",
	"BKS": "BKS",
	"DPK": "DPK",
}

var AdminFee = map[string]float64{
	"senin":  0.1,
	"rabu":   0.1,
	"jumat":  0.1,
	"selasa": 0.05,
	"kamis":  0.05,
	"sabtu":  0.05,
}

func CalculateAdminFee(price float32, day string) float32 {
	return price * float32(AdminFee[day])
}

var dayWithLocationCanDeliver = map[string][]string{
	"JKT": {"senin", "selasa", "rabu", "kamis", "jumat", "sabtu"},
	"BDG": {"rabu", "kamis", "sabtu"},
	"BKS": {"selasa", "kamis", "jumat"},
	"DPK": {"senin", "selasa"},
}

func IsDayCanDeliver(location string, day string) bool {
	for _, d := range dayWithLocationCanDeliver[location] {
		if d == day {
			return true
		}
	}
	return false
}

type Customer struct {
	FirstName string
	LastName  string
	Price     float32
	Location  string
}

type Customers []Customer

func CustomerFromString(data string) (Customer, error) {
	splitData := strings.Split(data, ":")

	locationCode, ok := LocationCode[splitData[3]]
	if !ok {
		return Customer{}, fmt.Errorf("Location code not found")
	}

	return Customer{
		FirstName: splitData[0],
		LastName:  splitData[1],
		Price:     StringToFloat32(splitData[2]),
		Location:  locationCode,
	}, nil
}

func StringToFloat32(str string) float32 {
	f, _ := strconv.ParseFloat(str, 32)
	return float32(f)
}

type DeliverOrderDetails struct {
	FullName   string
	TotalPrice float32
}

type DeliverOrders map[string]float32

func (DeliverOrders) Tranform(customers Customers) DeliverOrders {
	deliverOrders := DeliverOrders{}
	for _, c := range customers {
		deliverOrders[c.FirstName+"-"+c.LastName] += c.Price
	}
	return deliverOrders
}

func Deliver(customersData Customers, filter func(Customer) bool) Customers {
	customers := Customers{}
	for _, c := range customersData {
		if filter(c) {
			customers = append(customers, c)
		}
	}
	return customers
}

func CalculatePrice(customers Customers, day string) Customers {
	for i, v := range customers {
		customers[i].Price += CalculateAdminFee(v.Price, day)
	}
	return customers
}

func DeliveryOrder(data []string, day string) map[string]float32 {
	if len(data) == 0 {
		return map[string]float32{}
	}
	if day == "" {
		return map[string]float32{}
	}
	if _, ok := AdminFee[day]; !ok {
		return map[string]float32{}
	}

	customers := Customers{}
	for _, d := range data {
		customer, err := CustomerFromString(d)
		if err != nil {
			return map[string]float32{}
		}

		customers = append(customers, customer)
	}

	customers = Deliver(customers, func(c Customer) bool {
		return IsDayCanDeliver(c.Location, day)
	})

	customers = CalculatePrice(customers, day)

	deliverOrders := DeliverOrders{}.Tranform(customers)
	return deliverOrders
}

func main() {
	data := []string{
		"Budi:Gunawan:10000:JKTA",
		"Andi:Sukirman:20000:JKT",
		"Budi:Sukirman:30000:BDG",
		"Andi:Gunawan:40000:BKS",
		"Budi:Gunawan:50000:DPK",
	}

	day := "sabtu"

	deliveryData := DeliveryOrder(data, day)

	fmt.Println(deliveryData)
}
