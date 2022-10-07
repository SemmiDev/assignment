package model

import (
	"encoding/json"
	"time"
)

type Todo struct {
	Id   string `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

type Session struct {
	Username string
	Expiry   time.Time
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (er ErrorResponse) MarshalJSON() ([]byte, error) {
	return []byte(`{"error":"` + er.Error + `"}`), nil
}

//
//func (er ErrorResponse) UnmarshalJSON(data []byte) error {
//	err := json.Unmarshal(data, &er)
//	if err != nil {
//		return err
//	}
//	return nil
//}

type SuccessResponse struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func (sr SuccessResponse) MarshalJSON() ([]byte, error) {
	return []byte(`{"username":"` + sr.Username + `","message":"` + sr.Message + `"}`), nil
}

func MarshalJson(m json.Marshaler) ([]byte, error) {
	return m.MarshalJSON()
}
