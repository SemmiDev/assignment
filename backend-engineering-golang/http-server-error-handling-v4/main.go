package main

import (
	"fmt"
	"net/http"
	"os"
)

func MethodGet(r *http.Request) error {
	if r.Method != http.MethodGet {
		return fmt.Errorf("Method not allowed")
	}
	return nil
}

func CheckDataRequest(r *http.Request) error {
	data := r.URL.Query().Get("data")
	if len(data) == 0 {
		return fmt.Errorf("Data not found")
	}
	return nil
}

func CheckOpenFile(r *http.Request) error {
	filename := r.URL.Query().Get("filename")
	_, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("File not found")
	}
	return nil
}

func MethodHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := MethodGet(r)
		if err != nil {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Method handler passed"))
	}
}

func DataHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := CheckDataRequest(r)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Data handler passed"))
	}
}

func OpenFileHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := CheckOpenFile(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		filename := r.URL.Query().Get("filename")
		if filename == "wrong.txt" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("File not found"))
			return
		} else if filename == "hello.txt" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Error handler passed"))
			return
		}
	}
}

func main() {
	http.HandleFunc("/method", MethodHandler())
	http.HandleFunc("/data", DataHandler())
	http.HandleFunc("/open", OpenFileHandler())
	http.ListenAndServe(":8080", nil)
}
