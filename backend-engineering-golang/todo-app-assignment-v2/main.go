package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"

	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/middleware"
	"a21hc3NpZ25tZW50/model"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var credentials model.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errResponse, _ := model.MarshalJson(model.ErrorResponse{Error: "Internal Server Error"})
		w.Write(errResponse)
		return
	}

	if credentials.Username == "" || credentials.Password == "" {
		errResponse, _ := model.MarshalJson(model.ErrorResponse{Error: "Username or Password empty"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errResponse)
	}

	if db.Users[credentials.Username] != "" {
		errResponse, _ := model.MarshalJson(model.ErrorResponse{Error: "Username already exist"})
		w.WriteHeader(http.StatusConflict)
		w.Write(errResponse)
	}

	db.Users[credentials.Username] = credentials.Password
	w.WriteHeader(http.StatusOK)
	successResponse, _ := model.MarshalJson(model.SuccessResponse{Username: credentials.Username, Message: "Register Success"})
	w.Write(successResponse)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials model.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errResponse, _ := model.MarshalJson(model.ErrorResponse{Error: "Internal Server Error"})
		w.Write(errResponse)
		return
	}

	if credentials.Username == "" || credentials.Password == "" {
		errResponse, _ := model.MarshalJson(model.ErrorResponse{Error: "Username or Password empty"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errResponse)
	}

	if _, ok := db.Users[credentials.Username]; !ok {
		errResponse, _ := model.MarshalJson(model.ErrorResponse{Error: "Wrong User or Password!"})
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(errResponse)
	}

	if db.Users[credentials.Username] != credentials.Password {
		errResponse, _ := model.MarshalJson(model.ErrorResponse{Error: "Wrong User or Password!"})
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(errResponse)
	}

	session := model.Session{
		Username: credentials.Username,
		Expiry:   time.Now().Add(time.Hour * 5),
	}

	sessionKey := uuid.New().String()
	db.Sessions[sessionKey] = session

	cookie := http.Cookie{
		Name:    "session_token",
		Value:   sessionKey,
		Expires: session.Expiry,
	}
	http.SetCookie(w, &cookie)

	w.WriteHeader(http.StatusOK)
	successResponse, _ := model.MarshalJson(model.SuccessResponse{Username: credentials.Username, Message: "Login Success"})
	w.Write(successResponse)
}

func AddToDo(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errResponse, _ := model.MarshalJson(model.ErrorResponse{Error: "Internal Server Error"})
		w.Write(errResponse)
		return
	}

	if todo.Task == "" {
		errResponse, _ := model.MarshalJson(model.ErrorResponse{Error: "Name empty"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errResponse)
	}

	username := fmt.Sprintf("%s", r.Context().Value("username"))
	todo.Id = uuid.New().String()
	db.Task[username] = append(db.Task[username], todo)

	w.WriteHeader(http.StatusOK)
	msg := fmt.Sprintf("Task %s added!", todo.Task)
	successResponse, _ := model.MarshalJson(model.SuccessResponse{Username: username, Message: msg})
	w.Write(successResponse)
}

func ListToDo(w http.ResponseWriter, r *http.Request) {
	username := fmt.Sprintf("%s", r.Context().Value("username"))

	// jika task kosong, maka akan menampilkan pesan "Todolist not found!"
	// status code 404
	if len(db.Task) == 0 {
		w.WriteHeader(http.StatusNotFound)
		successResponse, _ := model.MarshalJson(model.ErrorResponse{Error: "Todolist not found!"})
		w.Write(successResponse)
		return
	}

	// menampilkan data dari map db.Task berdasarkan username
	// jika tidak ada data, makan cukup menampilkan status code 404
	tasks := db.Task[username]
	if len(tasks) == 0 {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusOK)
	tasksBytes, _ := json.Marshal(tasks)
	w.Write(tasksBytes)
}

func ClearToDo(w http.ResponseWriter, r *http.Request) {
	username := fmt.Sprintf("%s", r.Context().Value("username"))
	delete(db.Task, username)
	w.WriteHeader(http.StatusOK)
	successResponse, _ := model.MarshalJson(model.SuccessResponse{Username: username, Message: "Clear ToDo Success"})
	w.Write(successResponse)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	username := fmt.Sprintf("%s", r.Context().Value("username"))
	delete(db.Sessions, username)

	w.WriteHeader(http.StatusOK)
	successResponse, _ := model.MarshalJson(model.SuccessResponse{Username: username, Message: "Logout Success"})
	w.Write(successResponse)
}

func ResetToDo(w http.ResponseWriter, r *http.Request) {
	db.Task = map[string][]model.Todo{}
	w.WriteHeader(http.StatusOK)
}

type API struct {
	mux *http.ServeMux
}

func NewAPI() API {
	mux := http.NewServeMux()
	api := API{mux}

	mux.Handle("/user/register", middleware.Post(http.HandlerFunc(Register)))
	mux.Handle("/user/login", middleware.Post(http.HandlerFunc(Login)))
	mux.Handle("/user/logout", middleware.Get(middleware.Auth(http.HandlerFunc(Logout))))

	mux.Handle("/todo/create", middleware.Post(middleware.Auth(http.HandlerFunc(AddToDo))))
	mux.Handle("/todo/read", middleware.Get(middleware.Auth(http.HandlerFunc(ListToDo))))
	mux.Handle("/todo/clear", middleware.Delete(middleware.Auth(http.HandlerFunc(ClearToDo))))
	mux.Handle("/todo/reset", http.HandlerFunc(ResetToDo))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}

func main() {
	mainAPI := NewAPI()
	mainAPI.Start()
}
