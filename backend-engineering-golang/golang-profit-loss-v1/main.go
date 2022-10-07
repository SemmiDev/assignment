package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Transaction struct {
	Date   string
	Type   string
	Amount int
}

func formatTransaction(transaction Transaction) string {
	return fmt.Sprintf("%s;%s;%d", transaction.Date, transaction.Type, transaction.Amount)
}

func convertToInt(data string) int {
	dataInInt, _ := strconv.Atoi(data)
	return dataInInt
}

func extractToTransaction(data []string) []Transaction {
	transactions := make([]Transaction, 0, len(data))
	for _, d := range data {
		dSplit := strings.Split(d, ";")
		amount := convertToInt(dSplit[2])

		transaction := Transaction{
			Date:   dSplit[0],
			Type:   dSplit[1],
			Amount: amount,
		}

		transactions = append(transactions, transaction)
	}

	return transactions
}

func Readfile(path string) ([]string, error) {
	content, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer content.Close()

	fileScanner := bufio.NewScanner(content)
	fileScanner.Split(bufio.ScanLines)

	var extractedData []string
	for fileScanner.Scan() {
		extractedData = append(extractedData, fileScanner.Text())
	}

	// check if data is empty
	if extractedData == nil {
		return []string{}, nil
	}

	return extractedData, nil
}

func CalculateProfitLoss(data []string) string {
	transactions := extractToTransaction(data)
	lastTransaction := transactions[len(transactions)-1]

	totalProfit, totalLoss := 0, 0

	for _, transaction := range transactions {
		if transaction.Type == "income" {
			totalProfit += transaction.Amount
		} else {
			totalLoss += transaction.Amount
		}
	}

	if totalProfit == 0 {
		lastTransaction.Type = "loss"
		lastTransaction.Amount = totalLoss
		return formatTransaction(lastTransaction)
	}

	if totalLoss == 0 {
		lastTransaction.Type = "profit"
		lastTransaction.Amount = totalProfit
		return formatTransaction(lastTransaction)
	}

	if totalProfit > totalLoss {
		lastTransaction.Type = "profit"
		lastTransaction.Amount = totalProfit - totalLoss
		return formatTransaction(lastTransaction)
	}

	lastTransaction.Type = "loss"
	lastTransaction.Amount = totalLoss - totalProfit
	return formatTransaction(lastTransaction)
}

func main() {
	//data, err := Readfile("transactions.txt")
	//if err != nil {
	//	panic(err)
	//}

	datas := []string{
		"01/01/2021;income;1000",
		"02/01/2021;income;500",
		"03/01/2021;income;1000",
		"04/01/2021;income;1000",
		"05/01/2021;income;1000",
		"06/01/2021;income;1000",
		"07/01/2021;income;1000",
		"08/01/2021;income;1000",
		"09/01/2021;income;1000",
		"10/01/2021;income;1000",
	}
	profitLoss := CalculateProfitLoss(datas)
	fmt.Println(profitLoss)
}
