package main

import (
	"fmt"
	"strings"
)

func EmailInfo(email string) string {
	domain := strings.Split(strings.Split(email, "@")[1], ".")
	tld := strings.Join(domain[1:], ".")

	return fmt.Sprintf("Domain: %s dan TLD: %s", domain[0], tld)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("admin@yahoo.co.id"))
}
