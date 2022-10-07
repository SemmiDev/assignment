package invoice

import (
	"errors"
	"fmt"
	"time"
)

// Finance invoice
type FinanceInvoice struct {
	Date     string        // D/MM/YYYY (contoh: "01/01/2020")
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

var (
	ErrParseDate               = errors.New("error parse date")
	ErrInvoiceDetailsIsEmpty   = errors.New("invoice details is empty")
	ErrInvoiceStatusIsNotValid = errors.New("invoice status is not valid")
	ErrTotalPriceIsNotValid    = errors.New("total price is not valid")
)

func (fi FinanceInvoice) RecordInvoice() (InvoiceData, error) {
	if fi.Date == "" {
		return InvoiceData{}, ErrInvoiceDateIsEmpty
	}

	/*
		date = ("01/01/2020") to  ("01-January-2022")
	*/

	parseDate, err := time.Parse("02/01/2006", fi.Date)
	if err != nil {
		return InvoiceData{}, fmt.Errorf("%w: %v", ErrParseDate, err)
	}
	date := parseDate.Format("02-January-2006")

	if len(fi.Details) == 0 {
		return InvoiceData{}, ErrInvoiceDetailsIsEmpty
	}

	if fi.Status != PAID && fi.Status != UNPAID {
		return InvoiceData{}, ErrInvoiceStatusIsNotValid
	}

	totalInvoice := 0
	for _, detail := range fi.Details {
		if detail.Total == 0 {
			return InvoiceData{}, ErrTotalPriceIsNotValid
		} else {
			totalInvoice += detail.Total
		}
	}

	return InvoiceData{
		Date:         date,
		Departemen:   Finance,
		TotalInvoice: float64(totalInvoice),
	}, nil
}
