package invoice

import (
	"errors"
	"fmt"
	"strings"
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
	// if mi.Date == "" {
	// 	return InvoiceData{}, ErrInvoiceDateIsEmpty
	// }
	// if mi.StartDate == "" || mi.EndDate == "" {
	// 	return InvoiceData{}, ErrTravelDateIsEmpty
	// }
	// if mi.PricePerDay < 1 {
	// 	return InvoiceData{}, ErrPricePerDayIsNotValid
	// }

	// /*
	// 	date = ("01/01/2020") to  ("01-January-2022")
	// */

	// parseDate, err := time.Parse("02/01/2006", mi.Date)
	// if err != nil {
	// 	return InvoiceData{}, fmt.Errorf("%w: %v", ErrParseDate, err)
	// }
	// date := parseDate.Format("02-January-2006")

	// parseStartDate, err := time.Parse("02/01/2006", mi.StartDate)
	// if err != nil {
	// 	return InvoiceData{}, fmt.Errorf("%w: %v", ErrParseDate, err)
	// }
	// // startDate := parseStartDate.Format("02-January-2006")

	// parseEndDate, err := time.Parse("02/01/2006", mi.EndDate)
	// if err != nil {
	// 	return InvoiceData{}, fmt.Errorf("%w: %v", ErrParseDate, err)
	// }
	// // endDate := parseEndDate.Format("02-January-2006")

	// lengthOfTrip := parseEndDate.Sub(parseStartDate).Hours()/24 + 1
	// fmt.Println(lengthOfTrip)
	// totalInvoice := lengthOfTrip*float64(mi.PricePerDay) + float64(mi.AnotherFee)

	// return InvoiceData{
	// 	Date:         date,
	// 	TotalInvoice: totalInvoice,
	// 	Departemen:   Marketing,
	// }, nil

	invoice := InvoiceData{}
	var eror error

	//date
	if mi.Date == "" {
		eror = errors.New("invoice date is empty")
		return invoice, eror
	} else {
		invoice.Date = ChangeDate(mi.Date)
	}
	//start Date || End Date
	if mi.StartDate == "" || mi.EndDate == "" {
		eror = errors.New("travel date is empty")
		return invoice, eror
	}

	//Price per day
	if mi.PricePerDay <= 0 {
		eror = errors.New("price per day is not valid")
		return invoice, eror
	}

	//(end date - start date) x price per day + another fee
	splitStartDate := strings.Split(mi.StartDate, "/")
	splitEndDate := strings.Split(mi.EndDate, "/")
	StartDate, _ := time.Parse("2006 01 02", fmt.Sprintf("%s %s %s", splitStartDate[2], splitStartDate[1], splitStartDate[0]))
	EndDate, _ := time.Parse("2006 01 02", fmt.Sprintf("%s %s %s", splitEndDate[2], splitEndDate[1], splitEndDate[0]))

	hours := EndDate.Sub(StartDate).Hours()
	day := (int(hours) / 24)

	if !(splitStartDate[1] == "01" && splitEndDate[1] == "02" || splitStartDate[1] == "03" && splitEndDate[1] == "05") {
		day = day + 1
	}

	journey := float64(day) * float64(mi.PricePerDay)
	invoice.TotalInvoice = journey + float64(mi.AnotherFee)

	invoice.Departemen = Marketing

	return invoice, eror // TODO: replace this
}
