package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// Commaf produces a string form of the given number in base 10 with
// commas after every three orders of magnitude.
//
// e.g. Commaf(834142.32) -> 834,142.32
func Commaf(v float64) string {
	// Special case: zero.
	buf := &bytes.Buffer{}
	if v < 0 {
		// Special case: negative zero.
		buf.Write([]byte{'-'})
		v = 0 - v
	}

	// comma is the number of digits since the last comma.
	comma := []byte{','}

	// parts is the integer part of the number, in reverse order.
	parts := strings.Split(strconv.FormatFloat(v, 'f', -1, 64), ".")
	pos := 0
	if len(parts[0])%3 != 0 {
		pos += len(parts[0]) % 3
		buf.WriteString(parts[0][:pos])
		buf.Write(comma)
	}
	for ; pos < len(parts[0]); pos += 3 {
		buf.WriteString(parts[0][pos : pos+3])
		buf.Write(comma)
	}
	buf.Truncate(buf.Len() - 1)

	if len(parts) > 1 {
		buf.Write([]byte{'.'})
		buf.WriteString(parts[1])
	}
	return buf.String()
}

func stripTrailingDigits(s string, digits int) string {
	if i := strings.Index(s, "."); i >= 0 {
		if digits <= 0 {
			return s[:i]
		}
		i++
		if i+digits >= len(s) {
			return s
		}
		return s[:i+digits]
	}
	return s
}

func CommafWithDigits(f float64, decimals int) string {
	return stripTrailingDigits(Commaf(f), decimals)
}

func ChangeToCurrency(change int) string {
	humanizeValue := CommafWithDigits(float64(change), 0)
	stringValue := strings.Replace(humanizeValue, ",", ".", -1)
	return "Rp. " + stringValue

	// prefix := "Rp. "
	// // 10000 -> 10.000
	// // 100000 -> 100.000
	// // 1000000 -> 1.000.000
	// // 10000000 -> 10.000.000
	// return changer(prefix, change)
}

func MoneyChange(money int, listPrice ...int) string {
	totalPrice := 0
	for _, v := range listPrice {
		totalPrice += v
	}

	if money < totalPrice {
		return "Uang tidak cukup"
	}

	return ChangeToCurrency(money - totalPrice)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(MoneyChange(100000, 50000, 10000, 10000, 5000, 5000))
}
