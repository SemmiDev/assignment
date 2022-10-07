package main

type Product struct {
	Name  string
	Price int
	Tax   int
}

func loop(count int, value int, result *[]int) {
	for i := 1; i <= count; i++ {
		*result = append(*result, value)
	}
}

func ExchangeCoin(amount int) []int {
	banknotes := []int{1000, 500, 200, 100, 50, 20, 10, 5, 1}
	result := make([]int, 0, len(banknotes))
	for _, v := range banknotes {
		loop(amount/v, v, &result)
		amount = amount % v
	}
	return result
}

func MoneyChanges(amount int, products []Product) []int {
	var total int
	for _, v := range products {
		total += v.Price
		total += v.Tax
	}

	return ExchangeCoin(amount - total)
}
