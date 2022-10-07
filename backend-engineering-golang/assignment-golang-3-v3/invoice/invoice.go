package invoice

type InvoiceData struct {
	Date         string // "DD-month-YYYY" (contoh: "01-January-2022")
	TotalInvoice float64
	Departemen   DepartmentName
}

type DepartmentName string

func (d DepartmentName) ToString() string {
	switch d {
	case Finance:
		return "Finance"
	case Warehouse:
		return "Warehouse"
	case Marketing:
		return "Marketing"
	default:
		return "Unknown"
	}
}

const (
	Finance   DepartmentName = "finance"
	Warehouse DepartmentName = "warehouse"
	Marketing DepartmentName = "marketing"
)

type Invoice interface {
	RecordInvoice() (InvoiceData, error)
}

func generateKey(date string, department DepartmentName) string {
	return date + "#" + department.ToString()
}

func updateInvoicesData(invoicesData *[]InvoiceData, keyData string, data InvoiceData) {
	for i, invoiceData := range *invoicesData {
		key := generateKey(invoiceData.Date, invoiceData.Departemen)
		if keyData == key {
			(*invoicesData)[i] = data
		}
	}
}

func RecapDataInvoice(invoices []Invoice) ([]InvoiceData, error) {
	invoiceData := make([]InvoiceData, 0, len(invoices))
	sameDateAndDepartment := make(map[string]InvoiceData)

	for _, invoice := range invoices {
		switch invoice.(type) {
		case FinanceInvoice:
			financeInvoiceData, _ := invoice.(FinanceInvoice)
			if financeInvoiceData.Status == PAID && financeInvoiceData.Approved {
				financeRecordInvoice, err := financeInvoiceData.RecordInvoice()
				if err != nil {
					return nil, err
				}

				// date + department
				key := generateKey(financeRecordInvoice.Date, Finance)

				if _, ok := sameDateAndDepartment[key]; ok {
					totalInvoice := sameDateAndDepartment[key].TotalInvoice + financeRecordInvoice.TotalInvoice
					invoiceDataUpdated := InvoiceData{
						Date:         financeRecordInvoice.Date,
						TotalInvoice: totalInvoice,
						Departemen:   financeRecordInvoice.Departemen,
					}
					sameDateAndDepartment[key] = invoiceDataUpdated
					updateInvoicesData(&invoiceData, key, invoiceDataUpdated)
				} else {
					sameDateAndDepartment[key] = financeRecordInvoice
					invoiceData = append(invoiceData, financeRecordInvoice)
				}
			}
		case MarketingInvoice:
			marketingInvoiceData, _ := invoice.(MarketingInvoice)
			if marketingInvoiceData.Approved {
				marketingRecordInvoice, err := marketingInvoiceData.RecordInvoice()
				if err != nil {
					return nil, err
				}

				// date + department
				key := generateKey(marketingRecordInvoice.Date, Marketing)

				if existingData, ok := sameDateAndDepartment[key]; ok {
					totalInvoice := sameDateAndDepartment[key].TotalInvoice + marketingRecordInvoice.TotalInvoice
					invoiceDataUpdated := InvoiceData{
						Date:         existingData.Date,
						TotalInvoice: totalInvoice,
						Departemen:   existingData.Departemen,
					}
					sameDateAndDepartment[key] = invoiceDataUpdated
					updateInvoicesData(&invoiceData, key, invoiceDataUpdated)
				} else {
					sameDateAndDepartment[key] = marketingRecordInvoice
					invoiceData = append(invoiceData, marketingRecordInvoice)
				}
			}

		case WarehouseInvoice:
			warehouseInvoiceData, _ := invoice.(WarehouseInvoice)
			if warehouseInvoiceData.InvoiceType == PURCHASE && warehouseInvoiceData.Approved {
				warehouseRecordInvoice, err := warehouseInvoiceData.RecordInvoice()
				if err != nil {
					return nil, err
				}

				// date + department
				key := generateKey(warehouseRecordInvoice.Date, Warehouse)

				if _, ok := sameDateAndDepartment[key]; ok {
					totalInvoice := sameDateAndDepartment[key].TotalInvoice + warehouseRecordInvoice.TotalInvoice
					invoiceDataUpdated := InvoiceData{
						Date:         warehouseRecordInvoice.Date,
						TotalInvoice: totalInvoice,
						Departemen:   warehouseRecordInvoice.Departemen,
					}
					sameDateAndDepartment[key] = invoiceDataUpdated
					updateInvoicesData(&invoiceData, key, invoiceDataUpdated)
				} else {
					sameDateAndDepartment[key] = warehouseRecordInvoice
					invoiceData = append(invoiceData, warehouseRecordInvoice)
				}
			}
		}
	}

	return invoiceData, nil
}
