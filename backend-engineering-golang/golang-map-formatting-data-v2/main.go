package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)


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

	return getByKey(collections)
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


/*
	- Sam bikin struct yg punya field Header, Index, Position, dan Value
	- sam bikin struct yg punya field slice of struct yg di atas
	- Sam split slice data yg di input ke struct yg di atas
	- Sam bikin map yang ngeloop si collection of struct diatas, trus sam bkin key (header-index) dan value nya struct yg di atas
	-
*/

func getByKey(collections Collections) map[string][]string {
	datamap := make(map[string][]Data)
	for _, v := range collections.data {
		key := v.Header + "-" + v.Index
		datamap[key] = append(datamap[key], v)
	}

	newmap := make(map[string][]string)
	for k, v := range datamap {
		combine := combineTwoOneValue(v)
		newmap[k] = append(newmap[k], combine)
	}

	return grouping(newmap)
}

type Grouping struct {
	Key   string
	Index int
	Val   string
}

func grouping(data map[string][]string) map[string][]string {
	grouping := []Grouping{}
	for k, v := range data {
		key, index := strings.Split(k, "-")[0], strings.Split(k, "-")[1]
		indexInInt, _ := strconv.Atoi(index)

		group := Grouping{
			Key:   key,
			Index: indexInInt,
			Val:   v[0],
		}
		grouping = append(grouping, group)
	}

	newg := map[string][]Grouping{}
	for _, v := range grouping {
		key := v.Key
		newg[key] = append(newg[key], v)
	}

	for _, v := range newg {
		sort.Slice(v, func(i, j int) bool {
			return v[i].Index < v[j].Index
		})
	}

	newg2 := make(map[string][]string)
	for k, v := range newg {
		for _, v2 := range v {
			newg2[k] = append(newg2[k], v2.Val)
		}
	}
	return newg2
}

func combineTwoOneValue(data []Data) string {
	if len(data) == 1 {
		return data[0].Value
	}
	result := ""
	if data[0].Position == "first" {
		result = data[0].Value + " " + data[1].Value
	} else {
		result = data[1].Value + " " + data[0].Value
	}
	return result
}
