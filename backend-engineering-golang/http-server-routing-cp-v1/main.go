package main

import (
	"fmt"
	"net/http"
	"time"
)

func TimeHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		now := fmt.Sprintf("%v, %v %v %v", t.Weekday(), t.Day(), t.Month(), t.Year())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(now))
	})
}

func SayHelloHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello there"))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello, " + name + "!"))
		}
	})
}

func main() {
	http.HandleFunc("/time", TimeHandler())
	http.HandleFunc("/hello", SayHelloHandler())
	http.ListenAndServe("localhost:8080", nil)
}
