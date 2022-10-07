package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
input: data = ["account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe"]

output: map[account: ["John Doe", "Jane Doe"]]

input: data = ["account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe", "address-0-first-Jaksel", "address-0-last-Jakarta", "address-1-first-Bandung", "address-1-last-Jabar"]

output: map[account: ["John Doe", "Jane Doe"], address: ["Jaksel Jakarta", "Bandung Jabar"]]

header: account
index : 0
pos   :first
value: John

want map[account: ["John Doe", "Jane Doe"], address: ["Jaksel Jakarta", "Bandung Jabar"]]

	grouping based on header
		- account => [John Doe, Jane Doe]
		- address => [Jaksel Jakarta, Bandung Jabar]
	sorting based on index
		- account => [John Doe, Jane Doe]
		- address => [Jaksel Jakarta, Bandung Jabar]
	pair first and last
		- account => [John Doe, Jane Doe]
		- address => [Jaksel Jakarta, Bandung Jabar]

*/

func (c Collections) grouping() map[string][]string {
	// grouping based on header
	headers := make(map[string][]Data)
	for _, v := range c.data {
		headers[v.Header] = append(headers[v.Header], v)
	}

	// sorting based on index
	for _, v := range headers {
		sort.Slice(v, func(i, j int) bool {
			return v[i].Index < v[j].Index
		})
	}

	/*
		map[
			account:[{account 0 first John}   {account 0 last Doe}     {account 1 first Jane}    {account 1 last Doe}]
			address:[{address 0 first Jaksel} {address 0 last Jakarta} {address 1 first Bandung} {address 1 last Jabar}]
		]
	*/

	// pair first and last
	pairedMap := make(map[string][]string)
	for _, v := range headers {
		for i := 0; i < len(v); i += 2 {
			if v[i].Position == "first" && v[i+1].Position == "last" {
				pairedMap[v[i].Header] = append(pairedMap[v[i].Header], v[i].Value+" "+v[i+1].Value)
			}
		}
	}
	return pairedMap
}

type Data struct {
	Header   string
	Index    string
	Position string
	Value    string
}

type Collections struct {
	data []Data
}

func ChangeOutput(data []string) map[string][]string {
	var collections Collections

	for _, v := range data {
		split := strings.Split(v, "-")
		header := split[0]
		index := split[1]
		position := split[2]
		value := split[3]

		collections.data = append(collections.data, Data{
			Header:   header,
			Index:    index,
			Position: position,
			Value:    value,
		})
	}

	group := Group{
		collection: collections,
		data:       make(map[string]Data),
	}

	group.grouping()

	return nil
}

type Group struct {
	collection Collections

	// key = header + index
	// value = pos + value
	data map[string]Data
}

func (c *Group) grouping() {
	for _, v := range c.collection.data {
		key := v.Header + "-" + v.Index + "-" + v.Position
		c.data[key] = Data{
			Header:   v.Header,
			Index:    v.Index,
			Position: v.Position,
			Value:    v.Value,
		}
	}

	fmt.Println(c.data)
}

func main() {
	data := []string{
		"account-0-first-John",
		"account-0-last-Doe",
		"account-1-first-Jane",
		"account-1-last-Doe",
		"address-0-first-Jaksel",
		"address-0-last-Jakarta",
		"address-1-first-Bandung",
		"address-1-last-Jabar"}
	r := ChangeOutput(data)
	fmt.Println(r)
}
