package main

import (
	"context"
	"fmt"
	"time"

	"github.com/microgate-io/microgate-demo/todo"
	"github.com/microgate-io/microgate-demo/user"
	apilog "github.com/microgate-io/microgate-lib-go/v1/log"
	apiqueue "github.com/microgate-io/microgate-lib-go/v1/queue"
)

type TodoImpl struct {
	todo.UnimplementedTodoServiceServer

	// this is for calling a remote UserService
	userClient user.UserServiceClient

	// this is for notifying asynchronously about a Todo created
	queueClient apiqueue.QueueingServiceClient
}

func (s *TodoImpl) CreateTodo(ctx context.Context, req *todo.CreateTodoRequest) (*todo.CreateTodoResponse, error) {
	// check user using service

	userReq := &user.CheckUserRequest{
		Username: "lisa",
	}
	userResp, err := s.userClient.CheckUser(ctx, userReq)
	if err != nil {
		return nil, apilog.ErrorWithLog(ctx, err, "failed to check user", "username", "lisa", "title", req.Title)
	}
	if !userResp.IsValid {
		return nil, fmt.Errorf("user %s not valid", "lisa")
	}

	// store a ToDo in your database
	// simulate that by generating a new ID
	id := fmt.Sprintf("todo-%d", time.Now().UnixMilli())

	// publish
	_, err = s.queueClient.Publish(ctx, &apiqueue.PublishRequest{Topic: "todo", Message: []byte(req.Title)})
	if err != nil {
		return nil, apilog.ErrorWithLog(ctx, err, "failed to publish todo created")
	}

	return &todo.CreateTodoResponse{
		Id: id,
	}, nil
}
