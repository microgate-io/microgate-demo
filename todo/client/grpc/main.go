package main

import (
	"context"
	"log"

	"github.com/microgate-io/microgate-demo/todo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	opt := grpc.WithInsecure()
	addr := "localhost:9292"

	conn, err := grpc.Dial(addr, opt)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := todo.NewTodoServiceClient(conn)
	req := &todo.CreateTodoRequest{
		Title: "test",
	}
	withKey := metadata.AppendToOutgoingContext(context.Background(), "x-api-key", "goldenkey", "x-cloud-debug", "DEBUG")
	resp, err := client.CreateTodo(withKey, req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp.Id)
}
