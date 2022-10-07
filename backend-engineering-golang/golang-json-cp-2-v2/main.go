package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type LoanData struct {
	StartBalance int
	Data         []Loan
	Employees    []Employee
}

type Loan struct {
	Date        string
	EmployeeIDs []string
}

type Employee struct {
	ID       string
	Name     string
	Position string
}

type Borrower struct {
	ID        string `json:"id"`
	TotalLoan int    `json:"total_loan"`
}

type LoanRecord struct {
	MonthDate    string     `json:"month_date"`
	StartBalance int        `json:"start_balance"`
	EndBalance   int        `json:"end_balance"`
	Borrowers    []Borrower `json:"borrowers"`
}

func extractDate(date string) string {
	// 01-January-2021 -> January 2021
	dateSplit := strings.Split(date, "-")
	return fmt.Sprintf("%s %s", dateSplit[1], dateSplit[2])
}

func LoanReport(data LoanData) LoanRecord {
	loanRecord := LoanRecord{
		StartBalance: data.StartBalance,
		MonthDate:    extractDate(data.Data[0].Date),
	}

	elapsedBalance := 0
	loanMap := make(map[string]int)
	isDuidAbis := false
	for _, loan := range data.Data {
		for _, id := range loan.EmployeeIDs {
			if data.StartBalance-elapsedBalance <= 0 {
				fmt.Println(data.StartBalance - elapsedBalance)
				isDuidAbis = true
				break
			}
			loanMap[id] += 50_000
			elapsedBalance += 50_000
		}
	}

	if isDuidAbis {
		loanRecord.EndBalance = 0
	} else {
		loanRecord.EndBalance = data.StartBalance - elapsedBalance
	}

	for k, v := range loanMap {
		loanRecord.Borrowers = append(loanRecord.Borrowers,
			Borrower{
				ID:        k,
				TotalLoan: v,
			})
	}

	sort.Slice(loanRecord.Borrowers, func(i, j int) bool {
		return loanRecord.Borrowers[i].ID < loanRecord.Borrowers[j].ID
	})

	return loanRecord
}

func RecordJSON(record LoanRecord, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// encode json without new line in the end
	j, err := json.Marshal(record)
	if err != nil {
		return err
	}
	file.Write(j)
	return nil
}

// gunakan untuk debug
func main() {
	data := LoanData{
		StartBalance: 250000,
		Data: []Loan{
			{"01-March-2021", []string{"1", "2", "3", "4"}},
			{"04-March-2021", []string{"1", "2", "3"}},
		},
		Employees: []Employee{
			{"1", "Employee A", "Manager"},
			{"2", "Employee B", "Staff"},
			{"3", "Employee C", "Staff"},
			{"4", "Employee D", "Staff"},
		},
	}

	record := LoanReport(data)
	RecordJSON(record, "loan-records.json")

	//err := RecordJSON(records, "loan-records.json")
	//if err != nil {
	//	fmt.Println(err)
	//}

}

// 100000 * 2 = 200000 + 50000 = 250000
// 250000 + 50000 = 300000
