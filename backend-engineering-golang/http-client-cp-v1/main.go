package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Animechan struct {
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
}

func ClientGet() ([]Animechan, error) {
	client := http.Client{}
	listAnimeChan := make([]Animechan, 0, 10)

	// Hit API https://animechan.vercel.app/api/quotes/anime?title=naruto with method GET:
	req, _ := http.NewRequest("GET", "https://animechan.vercel.app/api/quotes/anime?title=naruto", nil)
	req.Header.Add("Content-Type", "application/json")
	res, _ := client.Do(req)

	// parse the response body
	body, _ := io.ReadAll(res.Body)
	err := json.Unmarshal(body, &listAnimeChan)
	if err != nil {
		return nil, err
	}
	return listAnimeChan, nil
}

type data struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Postman struct {
	Data data
	Url  string `json:"url"`
}

func ClientPost() (Postman, error) {
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Dion",
		"email": "dionbe2022@gmail.com",
	})
	responseBody := bytes.NewBuffer(postBody)
	fmt.Println(responseBody)

	// Hit API https://postman-echo.com/post with method POST:
	resp, _ := http.Post("https://postman-echo.com/post", "application/json", responseBody)
	body, _ := ioutil.ReadAll(resp.Body)
	var postman Postman
	err := json.Unmarshal(body, &postman)
	if err != nil {
		return Postman{}, err
	}
	return postman, nil
}

func main() {
	get, _ := ClientGet()
	fmt.Println(get)

	post, _ := ClientPost()
	fmt.Println(post)
}
