package main

import (
	"fmt"
	"net/http"
	"time"
)

func GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// expected output =  Thursday, 22 September 2022
		t := time.Now()
		formatted := fmt.Sprintf("%v, %v %v %v", t.Weekday(), t.Day(), t.Month(), t.Year())
		writer.Write([]byte(formatted))
	}
}

func main() {
	http.ListenAndServe("localhost:8080", GetHandler())
}
