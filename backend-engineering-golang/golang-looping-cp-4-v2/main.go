package main

import (
	"fmt"
	"strings"
)

func EmailInfo(email string) string {
	info := strings.Split(email, "@")[1]
	info2 := strings.Split(info, ".")
	info3 := strings.Join(info2[1:], ".")

	return fmt.Sprintf("Domain: %s dan TLD: %s", info2[0], info3)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("admin@yahoo.co.id"))
}
