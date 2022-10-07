package main

import (
	"a21hc3NpZ25tZW50/invoice"
	"fmt"
	"log"
)

func RecapDataInvoice(data []invoice.Invoice) ([]invoice.InvoiceData, error) {
	return invoice.RecapDataInvoice(data)
}

func main() {
	listInvoice := []invoice.Invoice{
		invoice.FinanceInvoice{
			Date:     "01/02/2020",
			Details:  []invoice.Detail{{Description: "pembelian nota", Total: 4000}, {Description: "Pembelian alat tulis", Total: 4000}},
			Status:   invoice.PAID,
			Approved: true,
		},
		invoice.FinanceInvoice{
			Date:     "01/02/2020",
			Details:  []invoice.Detail{{Description: "pembelian nota", Total: 4000}, {Description: "Pembelian alat tulis", Total: 4000}},
			Status:   invoice.PAID,
			Approved: true,
		},
		invoice.WarehouseInvoice{
			Date: "01-February-2020",
			Products: []invoice.Product{
				{Name: "product A", Unit: 10, Price: 10000, Discount: 0.1},
				{Name: "product C", Unit: 5, Price: 15000, Discount: 0.2},
			},
			InvoiceType: invoice.PURCHASE,
			Approved:    true,
		},
		invoice.MarketingInvoice{
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

	fmt.Println("-----------")
	fmt.Println(result)
	expected := []invoice.InvoiceData{
		{Date: "01-February-2020", TotalInvoice: 16000, Departemen: invoice.Finance},
		{Date: "01-February-2020", TotalInvoice: 65000, Departemen: invoice.Marketing},
		{Date: "01-February-2020", TotalInvoice: 150000, Departemen: invoice.Warehouse},
	}
	fmt.Println(expected)
	fmt.Println("-----------")

	fmt.Println(compare(result, expected))
}

func compare(a, b []invoice.InvoiceData) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
