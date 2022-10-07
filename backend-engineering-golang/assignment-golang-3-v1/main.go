package main

import (
	"fmt"
	"log"
)

type Invoice interface {
	RecordInvoice() (InvoiceData, error)
}


// Finance invoice
type FinanceInvoice struct {
	Date     string
	Status   InvoiceStatus // status: "paid", "unpaid"
	Approved bool
	Details  []Detail
}

type InvoiceStatus string

const (
	PAID   InvoiceStatus = "paid"
	UNPAID InvoiceStatus = "unpaid"
)

type Detail struct {
	Description string
	Total       int
}


// Warehouse invoice

type WarehouseInvoice struct {
	Date        string
	InvoiceType InvoiceTypeName
	Approved    bool
	Products    []Product
}

type InvoiceTypeName string

const (
	PURCHASE InvoiceTypeName = "purchase"
	SALES    InvoiceTypeName = "sales"
)

type Product struct {
	Name     string
	Unit     int
	Price    int
	Discount float64
}


// Marketing invoice
type MarketingInvoice struct {
	Date        string
	StartDate   string
	EndDate     string
	PricePerDay int
	AnotherFee  int
	Approved    bool
}


type InvoiceData struct {
	Date         string
	TotalInvoice float64
	Departemen   DepartmentName
}

type DepartmentName string

const (
	Finance   DepartmentName = "finance"
	Warehouse DepartmentName = "warehouse"
	Marketing DepartmentName = "marketing"
)


func RecapDataInvoice(data []Invoice) ([]InvoiceData, error) {
	// TODO: answer here
}

func main() {
	listInvoice := []Invoice{
		FinanceInvoice{
			Date:     "01/02/2020",
			Details:  []Detail{{"pembelian nota", 4000}, {"Pembelian alat tulis", 4000}},
			Status:   PAID,
			Approved: true,
		},
		FinanceInvoice{
			Date:     "01/02/2020",
			Details:  []Detail{{"pembelian nota", 4000}, {"Pembelian alat tulis", 4000}},
			Status:   PAID,
			Approved: true,
		},
		WarehouseInvoice{
			Date: "01-February-2020",
			Products: []Product{
				{"product A", 10, 10000, 0.1},
				{"product C", 5, 15000, 0.2},
			},
			InvoiceType: PURCHASE,
			Approved:    true,
		},
		MarketingInvoice{
			Date:        "01/02/2020",
			StartDate:   "20/01/2020",
			EndDate:     "25/01/2020",
			Approved:    true,
			PricePerDay: 10000,
			AnotherFee:  5000,
		},
	}

	result, err := RecapDataInvoice(listInvoice)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
