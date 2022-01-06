package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	body := strings.NewReader(`{"title":"hello microgate"}`)
	resp, err := http.Post("http://localhost:8080/todo/v1/todo-service/create-todo", "application/json", body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp.Status)
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(data))
}
