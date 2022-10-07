package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"
)

func Get(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			errResponse, _ := model.MarshalJson(model.ErrorResponse{Error: "Method is not allowed!"})
			w.Write(errResponse)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Post(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			errResponse, _ := model.MarshalJson(model.ErrorResponse{Error: "Method is not allowed!"})
			w.Write(errResponse)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Delete(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			errResponse, _ := model.MarshalJson(model.ErrorResponse{Error: "Method is not allowed!"})
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write(errResponse)
			return
		}
		next.ServeHTTP(w, r)
	})
}
