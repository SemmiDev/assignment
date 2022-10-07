package main

import (
	"math/rand"
	"net/http"
	"time"
)

var Quotes = []string{
	"Be yourself; everyone else is already taken. ― Oscar Wilde",
	"Be the change that you wish to see in the world. ― Mahatma Gandhi",
	"I have not failed. I've just found 10,000 ways that won't work. ― Thomas A. Edison",
	"It is never too late to be what you might have been. ― George Eliot",
	"Everything you can imagine is real. ― Pablo Picasso",
	"Nothing is impossible, the word itself says 'I'm possible'! ― Audrey Hepburn",
}

type QuotesHandler struct{}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomQuotes(quotes []string) string {
	return quotes[rand.Intn(len(quotes))]
}

func (qh QuotesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(randomQuotes(Quotes)))
}

func main() {
	handler := QuotesHandler{}
	http.ListenAndServe("localhost:8080", handler)
}
