package invoice

import (
	"errors"
)

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
	Price    float64
	Discount float64
}

var (
	ErrInvoiceDateIsEmpty     = errors.New("invoice date is empty")
	ErrInvoiceProductsIsEmpty = errors.New("invoice products is empty")
	ErrInvoiceTypeIsNotValid  = errors.New("invoice type is not valid")
	ErrUnitProductIsNotValid  = errors.New("unit product is not valid")
	ErrPriceProductIsNotValid = errors.New("price product is not valid")
)

func (wi WarehouseInvoice) RecordInvoice() (InvoiceData, error) {
	if wi.Date == "" {
		return InvoiceData{}, ErrInvoiceDateIsEmpty
	}

	if wi.InvoiceType != PURCHASE && wi.InvoiceType != SALES {
		return InvoiceData{}, ErrInvoiceTypeIsNotValid
	}

	if len(wi.Products) == 0 {
		return InvoiceData{}, ErrInvoiceProductsIsEmpty
	}

	for _, product := range wi.Products {
		if product.Unit < 1 {
			return InvoiceData{}, ErrUnitProductIsNotValid
		}
		if product.Price < 1 {
			return InvoiceData{}, ErrPriceProductIsNotValid
		}
	}

	var totalInvoice float64
	for _, product := range wi.Products {
		total := product.Price * float64(product.Unit)
		discount := total * product.Discount
		totalWithDiscount := total - discount
		totalInvoice += totalWithDiscount
	}

	return InvoiceData{
		Date:         wi.Date,
		TotalInvoice: totalInvoice,
		Departemen:   Warehouse,
	}, nil
}
