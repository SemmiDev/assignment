package middleware

import (
	"a21hc3NpZ25tZW50/db"
	"context"
	"net/http"
	"time"

	"a21hc3NpZ25tZW50/model"
)

func isExpired(s model.Session) bool {
	return s.Expiry.Before(time.Now())
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := r.Cookie("session_token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			errResponse, _ := model.MarshalJson(model.ErrorResponse{Error: "http: named cookie not present"})
			w.Write(errResponse)
			return
		}

		s, ok := db.Sessions[session.Value]
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			errResponse, _ := model.MarshalJson(model.ErrorResponse{Error: "Session not found"})
			w.Write(errResponse)
			return
		}

		if isExpired(s) {
			w.WriteHeader(http.StatusUnauthorized)
			errResponse, _ := model.MarshalJson(model.ErrorResponse{Error: "Session expired"})
			w.Write(errResponse)
			return
		}

		// set username to context
		ctx := r.Context()
		ctx = context.WithValue(ctx, "username", s.Username)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
