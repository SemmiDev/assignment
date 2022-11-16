package main

import (
	"fmt"
	"os"
	"strings"
)

type Transaction struct {
	Date   string
	Type   string
	Amount int
}

func (t Transaction) format() string {
	return fmt.Sprintf("%s;%s;%d", t.Date, t.Type, t.Amount)
}

func generateKey(date string, transactionType string) string {
	return fmt.Sprintf("%s-%s", date, transactionType)
}

func resolveKey(key string) (string, string) {
	split := strings.Split(key, "-")
	return split[0], split[1]
}

func sortTransactionsByDay(transactions []Transaction) []Transaction {
	for i := 0; i < len(transactions); i++ {
		for j := i + 1; j < len(transactions); j++ {
			if transactions[i].Date > transactions[j].Date {
				transactions[i], transactions[j] = transactions[j], transactions[i]
			}
		}
	}
	return transactions
}

/*

	- grouping hari yg sama
*/

func GroupingSameDay(transaction []Transaction) []Transaction {
	sameDay := map[string]int{}
	/*
		key = 01/01/2021-income
	*/
	for _, v := range transaction {
		key := generateKey(v.Date, v.Type)
		sameDay[key] += v.Amount
	}

	sameDay = Accumulate(sameDay)

	data := make([]Transaction, 0, len(sameDay))
	for k, v := range sameDay {
		date, transactionType := resolveKey(k)
		t := Transaction{
			Date:   date,
			Type:   transactionType,
			Amount: v,
		}
		data = append(data, t)
	}
	data = sortTransactionsByDay(data)
	return data
}

func Accumulate(transactions map[string]int) map[string]int {
	maps := map[string][]Transaction{}

	for k, v := range transactions {
		date, txType := resolveKey(k)
		maps[date] = append(maps[date], Transaction{
			Date:   date,
			Type:   txType,
			Amount: v,
		})
	}

	newTransactions := map[string]int{}
	for k, v := range maps {
		// k is date

		if len(v) == 1 {
			if v[0].Type == "income" {
				key := generateKey(k, "income")
				newTransactions[key] = v[0].Amount
			}
			if v[0].Type == "expense" {
				key := generateKey(k, "expense")
				newTransactions[key] = v[0].Amount
			}

			continue
		}

		if len(v) == 2 {
			if v[0].Type == "income" && v[1].Type == "expense" {
				acc := v[0].Amount - v[1].Amount
				if v[0].Amount > v[1].Amount {
					key := generateKey(k, "income")
					newTransactions[key] = acc
				} else {
					key := generateKey(k, "expense")
					newTransactions[key] = -acc
				}
			} else if v[0].Type == "expense" && v[1].Type == "income" {
				acc := v[1].Amount - v[0].Amount
				if v[1].Amount > v[0].Amount {
					key := generateKey(k, "income")
					newTransactions[key] = acc
				} else {
					key := generateKey(k, "expense")
					newTransactions[key] = -acc
				}
			}
		}
	}

	return newTransactions
}

func RecordTransactions(path string, transactions []Transaction) error {
	if len(transactions) == 0 {
		return nil
	}

	groupingSameDay := GroupingSameDay(transactions)

	// remove transactions.txt if exists
	if _, err := os.Stat(path); err == nil {
		err := os.Remove(path)
		if err != nil {
			return err
		}
	}

	// open or create transactions.txt
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// write to file
	for i, v := range groupingSameDay {
		// if last data don't add new line
		if i == len(groupingSameDay)-1 {
			file.WriteString(v.format())
		} else {
			file.WriteString(v.format() + "\n")
		}
	}

	//totalIncome, totalExpense := totalIncomeExpense(transactions)
	return nil
	//return errors.New("not implemented") // TODO: replace this
}

func main() {

	// 200000 + 1000 = 201000 (income)
	// 100000 + 1000 = 101000 (expense)
	// akumulasi = 201000 - 101000 = 100000 (income)

	// 20000 + 233000 = 253000 (income)

	// bisa digunakan untuk pengujian test case
	var transactions = []Transaction{
		{"01/01/2021", "income", 200000},
		{"01/01/2021", "income", 1000},
		{"01/01/2021", "expense", 100000},
		{"01/01/2021", "expense", 1000},
		{"02/01/2021", "income", 20000},
		{"02/01/2021", "income", 233000},
		{"02/01/2021", "expense", 3424},
		{"02/01/2021", "expense", 2300},
		{"02/01/2021", "expense", 42000},
		{"03/01/2021", "income", 20000},
		{"03/01/2021", "income", 22000},
		{"03/01/2021", "expense", 22321},
		{"04/01/2021", "income", 24000},
		{"04/01/2021", "income", 20000},
		{"04/01/2021", "expense", 223200},
		{"05/01/2021", "income", 20000},
		{"05/01/2021", "income", 50000},
		{"05/01/2021", "expense", 2213},
		{"06/01/2021", "income", 60000},
		{"06/01/2021", "income", 70000},
		{"06/01/2021", "expense", 4545},
		{"07/01/2021", "income", 80000},
		{"07/01/2021", "income", 110000},
		{"07/01/2021", "income", 120000},
		{"07/01/2021", "income", 111200},
		{"07/01/2021", "expense", 55500},
		{"08/01/2021", "expense", 200000},
		{"10/01/2021", "expense", 20000},
		{"11/01/2021", "expense", 10000},
		{"12/01/2021", "expense", 55500},
		{"13/01/2021", "expense", 55500},
		{"02/01/2021", "expense", 55500},
		{"02/01/2021", "expense", 10000},
		{"14/01/2021", "expense", 20000},
		{"11/01/2021", "expense", 20000},
		{"15/01/2021", "expense", 10000},
		{"16/01/2021", "expense", 20000},
		{"02/01/2021", "expense", 55500},
		{"17/01/2021", "expense", 10000},
		{"06/01/2021", "expense", 20000},
		{"18/01/2021", "expense", 10000},
		{"03/01/2021", "expense", 20000},
		{"04/01/2021", "expense", 10000},
		{"19/01/2021", "expense", 55500},
		{"20/01/2021", "expense", 55500},
		{"21/01/2021", "expense", 10000},
		{"22/01/2021", "expense", 10000},
		{"23/01/2021", "expense", 10000},
		{"24/01/2021", "expense", 10000},
	}

	err := RecordTransactions("transactions.txt", transactions)
	if err != nil {
		panic(err)
	}

	//fmt.Println("Success")
}
