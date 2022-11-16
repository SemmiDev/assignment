package api

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"net/http"
	"path"
	"text/template"
	"time"

	"github.com/google/uuid"
)

func (api *API) Register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	creds := model.Credentials{
		Username: r.Form.Get("username"),
		Password: r.Form.Get("password"),
	}

	if creds.Username == "" || creds.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Username or Password empty"})
		return
	}

	err = api.usersRepo.AddUser(creds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	filepath := path.Join("views", "status.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	var data = map[string]string{"name": creds.Username, "message": "register success!"}
	err = tmpl.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
	}
}

func (api *API) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	creds := model.Credentials{
		Username: r.Form.Get("username"),
		Password: r.Form.Get("password"),
	}

	if creds.Username == "" || creds.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Username or Password empty"})
		return
	}

	err = api.usersRepo.LoginValid(creds)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	sessionKey := uuid.New().String()
	expiry := time.Now().Local().Add(time.Hour * 5)

	cookie := http.Cookie{
		Path:    "/",
		Name:    "session_token",
		Value:   sessionKey,
		Expires: expiry,
	}

	session := model.Session{
		Token:    sessionKey,
		Username: creds.Username,
		Expiry:   expiry,
	}

	err = api.sessionsRepo.AddSessions(session)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	http.SetCookie(w, &cookie)
	api.dashboardView(w, r)
}

func (api *API) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	sessionToken := cookie.Value
	api.sessionsRepo.DeleteSessions(sessionToken)

	cookie = &http.Cookie{
		Path:    "/",
		Name:    "session_token",
		Value:   "",
		Expires: time.Now(),
	}

	http.SetCookie(w, cookie)

	filepath := path.Join("views", "login.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
	}
}
