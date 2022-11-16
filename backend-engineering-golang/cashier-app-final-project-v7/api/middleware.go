package api

import (
	"a21hc3NpZ25tZW50/model"
	"context"
	"encoding/json"
	"net/http"
)

func (api *API) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionToken, err := r.Cookie("session_token") // TODO: replace this
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
			return
		}

		sessionValue := sessionToken.Value
		sessionFound, err := api.sessionsRepo.CheckExpireToken(sessionValue)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), "username", sessionFound.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (api *API) Get(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Method is not allowed!"})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (api *API) Post(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Method is not allowed!"})
			return
		}
		next.ServeHTTP(w, r)
	})
}
