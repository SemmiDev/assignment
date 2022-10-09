package api

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"path"
	"text/template"
	"time"
)

func (api *API) Register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	creds := model.Credentials{
		Username: username,
		Password: password,
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
		_ = json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	creeds := model.Credentials{
		Username: username,
		Password: password,
	}

	if creeds.Username == "" || creeds.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Username or Password empty"})
		return
	}

	listUser, err := api.usersRepo.ReadUser()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	if !api.usersRepo.LoginValid(listUser, creeds) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Wrong User or Password!"})
		return
	}

	// Generate Cookie with Name "session_token", Path "/", Value "uuid generated with github.com/google/uuid", Expires time to 5 Hour.
	sessionKey := uuid.New().String()
	expiry := time.Now().Local().Add(time.Hour * 5)

	cookie := http.Cookie{
		Path:    "/",
		Name:    "session_token",
		Value:   sessionKey,
		Expires: expiry,
	}
	http.SetCookie(w, &cookie)

	session := model.Session{
		Token:    sessionKey,
		Username: creeds.Username,
		Expiry:   expiry,
	}
	err = api.sessionsRepo.AddSessions(session)

	filepath := path.Join("views", "dashboard.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	products, _ := api.products.ReadProducts()
	carts, _ := api.cartsRepo.ReadCart()

	// sesuiai yg login
	carts.Name = creeds.Username

	data := model.Dashboard{
		Product: products,
		Cart:    carts,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
	}
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
