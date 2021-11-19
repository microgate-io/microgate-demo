package main

import (
	"context"
	"log"

	"github.com/shurcooL/graphql"
)

type CreateTodoRequest struct {
}

func main() {
	client := graphql.NewClient("http://localhost:8080/graphql", nil)
	query := CreateTodoRequest{}
	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		log.Println(err)
	}

}
