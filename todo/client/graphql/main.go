package main

import (
	"context"
	"log"

	"github.com/Khan/genqlient/graphql"
)

func main() {
	client := graphql.NewClient("http://localhost:8080/query", nil)
	req := CreateTodoRequestInput{Title: "dennis"}
	resp, err := todoServiceCreateTodo(context.Background(), client, req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.TodoServiceCreateTodo.Id)
}

//go:generate go run github.com/Khan/genqlient genqlient.yaml
