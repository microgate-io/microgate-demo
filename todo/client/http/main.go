package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	body := strings.NewReader(`
{}	
	`)
	resp, err := http.Post("http://localhost:8080/main/TodoService/CreateTodo", "application/json", body)
	if err != nil {
		log.Fatalln(err)
	}
	data, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(data))
}
