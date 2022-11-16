package main

import (
	"errors"
	"strings"
	"fmt"
	tld "github.com/jpillora/go-tld"
)

var mappedDomainToIDNTLD = map[string]string{
	".com": ".co.id",
	".org": ".org.id",
	".gov": ".go.id",
}

var IDNDomain = map[string]string{
	"com": ".co.id",
	"org": ".org.id",
	"gov": ".go.id",
}

type RowData struct {
	RankWebsite int
	Domain      string
	TLD         string
	IDN_TLD     string
	Valid       bool
	RefIPs      int
}

func ProcessGetTLD(website RowData, ch chan RowData, chErr chan error) {
	if website.Domain == "" {
		chErr <- errors.New("domain name is empty")
	}
	if !website.Valid {
		chErr <- errors.New("domain not valid")
	}
	if website.RefIPs < 0 {
		chErr <- errors.New("domain RefIPs not valid")
	}

	split := strings.Split(website.Domain, ".")
	domainSplit := strings.Join(split[1:], ".")

	if exist, ok := IDNDomain[domainSplit]; ok {
		website.TLD = "." + domainSplit
		website.IDN_TLD = exist
	} else {
		website.TLD = "." + domainSplit
		website.IDN_TLD = domainSplit
	}

	ch <- website
}

// func ProcessGetTLD(website RowData, ch chan RowData, chErr chan error) {
// 	if website.Domain == "" {
// 		chErr <- errors.New("domain name is empty")
// 		return
// 	}

// 	if !website.Valid {
// 		chErr <- errors.New("domain not valid")
// 		return
// 	}

// 	if website.RefIPs < 0 {
// 		chErr <- errors.New("domain RefIPs not valid")
// 		return
// 	}

// 	prefix := "https://" + website.Domain
// 	parser, err := tld.Parse(prefix)
// 	if err != nil {
// 		chErr <- err
// 		return
// 	}

// 	tld := "." + parser.TLD
// 	website.TLD = tld
// 	website.IDN_TLD = tld

// 	if idnTLD, ok := mappedDomainToIDNTLD[tld]; ok {
// 		website.IDN_TLD = idnTLD
// 	}

// 	ch <- website
// }

var FuncProcessGetTLD = ProcessGetTLD

func FilterAndFillData(TLD string, data []RowData) ([]RowData, error) {
	ch := make(chan RowData, len(data)) // buffered channel = [,,, len(data)]
	errCh := make(chan error) // unbuffer = 0


	for _, website := range data {
		go FuncProcessGetTLD(website, ch, errCh)
	}


	var result []RowData
	
	for i := 0; i < len(data); i++ {
		select {
		case website := <-ch:
			if website.TLD == TLD {
				result = append(result, website)
			}
		case err := <-errCh:
			return nil, err
		}
	}

	return result, nil
}

func main() {
	// fck wkwkw
	parser, err := tld.Parse("https://google.co.id")
	fmt.Println(err)
	fmt.Println(parser.Domain)
	fmt.Println(parser.TLD)
}
