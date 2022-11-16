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
	// if wi.Date == "" {
	// 	return InvoiceData{}, ErrInvoiceDateIsEmpty
	// }

	// if wi.InvoiceType != PURCHASE && wi.InvoiceType != SALES {
	// 	return InvoiceData{}, ErrInvoiceTypeIsNotValid
	// }

	// if len(wi.Products) == 0 {
	// 	return InvoiceData{}, ErrInvoiceProductsIsEmpty
	// }

	// for _, product := range wi.Products {
	// 	if product.Unit < 1 {
	// 		return InvoiceData{}, ErrUnitProductIsNotValid
	// 	}
	// 	if product.Price < 1 {
	// 		return InvoiceData{}, ErrPriceProductIsNotValid
	// 	}
	// }

	// var totalInvoice float64
	// for _, product := range wi.Products {
	// 	total := product.Price * float64(product.Unit)
	// 	discount := total * product.Discount
	// 	totalWithDiscount := total - discount
	// 	totalInvoice += totalWithDiscount
	// }

	// return InvoiceData{
	// 	Date:         wi.Date,
	// 	TotalInvoice: totalInvoice,
	// 	Departemen:   Warehouse,
	// }, nil

	invoice := InvoiceData{}
	var eror error

	//date
	if wi.Date == "" {
		eror = errors.New("invoice date is empty")
		return invoice, eror
	} else {
		invoice.Date = ChangeDate(wi.Date)
	}

	if wi.InvoiceType != PURCHASE && wi.InvoiceType != SALES {

		eror = errors.New("invoice type is invalid")
		return invoice, eror
	}

	//invoice type
	if wi.InvoiceType == "" {
		eror = errors.New("invoice type is invalid")
		return invoice, eror
	}

	//products
	if len(wi.Products) == 0 {
		eror = errors.New("invoice products is empty")
		return invoice, eror
	}

	//unit
	var total float64
	for _, v := range wi.Products {
		if v.Unit <= 0 {
			eror = errors.New("unit product is not valid")
			return invoice, eror
		}
		if v.Price <= 0 {
			eror = errors.New("price product is not valid")
			return invoice, eror
		}
		var discount float64
		product := (float64(v.Unit) * v.Price)
		discount = product * v.Discount
		total = product - discount
	}

	for _, v := range wi.Products {
		var discount float64
		product := (float64(v.Unit) * v.Price)
		discount = product * v.Discount
		total = product - discount
		invoice.TotalInvoice = total + invoice.TotalInvoice
	}
	invoice.Departemen = Warehouse
	return invoice, eror // TODO: replace this
}
