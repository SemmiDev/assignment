package invoice

import (
	"errors"
	"fmt"
	"time"
)

// Marketing invoice
type MarketingInvoice struct {
	Date        string
	StartDate   string
	EndDate     string
	PricePerDay int
	AnotherFee  int
	Approved    bool
}

var (
	ErrTravelDateIsEmpty     = errors.New("travel date is empty")
	ErrPricePerDayIsNotValid = errors.New("price per day is not valid")
)

func (mi MarketingInvoice) RecordInvoice() (InvoiceData, error) {
	if mi.Date == "" {
		return InvoiceData{}, ErrInvoiceDateIsEmpty
	}
	if mi.StartDate == "" || mi.EndDate == "" {
		return InvoiceData{}, ErrTravelDateIsEmpty
	}
	if mi.PricePerDay < 1 {
		return InvoiceData{}, ErrPricePerDayIsNotValid
	}

	/*
		date = ("01/01/2020") to  ("01-January-2022")
	*/

	parseDate, err := time.Parse("02/01/2006", mi.Date)
	if err != nil {
		return InvoiceData{}, fmt.Errorf("%w: %v", ErrParseDate, err)
	}
	date := parseDate.Format("02-January-2006")

	parseStartDate, err := time.Parse("02/01/2006", mi.StartDate)
	if err != nil {
		return InvoiceData{}, fmt.Errorf("%w: %v", ErrParseDate, err)
	}
	// startDate := parseStartDate.Format("02-January-2006")

	parseEndDate, err := time.Parse("02/01/2006", mi.EndDate)
	if err != nil {
		return InvoiceData{}, fmt.Errorf("%w: %v", ErrParseDate, err)
	}
	// endDate := parseEndDate.Format("02-January-2006")

	lengthOfTrip := parseEndDate.Sub(parseStartDate).Hours()/24 + 1
	fmt.Println(lengthOfTrip)
	totalInvoice := lengthOfTrip*float64(mi.PricePerDay) + float64(mi.AnotherFee)

	return InvoiceData{
		Date:         date,
		TotalInvoice: totalInvoice,
		Departemen:   Marketing,
	}, nil
}
