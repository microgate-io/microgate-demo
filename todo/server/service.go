package main

import (
	"context"
	"fmt"
	"time"

	"github.com/microgate-io/microgate-demo/todo"
	"github.com/microgate-io/microgate-demo/user"
	apidb "github.com/microgate-io/microgate-lib-go/v1/db"
	apilog "github.com/microgate-io/microgate-lib-go/v1/log"
	apiqueue "github.com/microgate-io/microgate-lib-go/v1/queue"
)

type TodoImpl struct {
	todo.UnimplementedTodoServiceServer

	// this is for calling a remote UserService
	userClient user.UserServiceClient

	// this is for calling a local Microgate DatabaseService
	dbClient apidb.DatabaseServiceClient

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

	var id string
	apidb.InTransactionDo(ctx, s.dbClient, func(transaction_id string) error {
		mResp, err := s.dbClient.Mutate(ctx, &apidb.MutationRequest{
			TransactionId: transaction_id,
			Sql:           fmt.Sprintf("INSERT INTO public.tasks (title) VALUES ('hello microgate %v') RETURNING task_id", time.Now())})

		// TEMP hack
		apilog.Debugw(ctx, "insert result", "rows", mResp.Rows)
		if mResp.Error == "" {
			if len(mResp.Rows) > 0 {
				if len(mResp.Rows[0].Columns) > 0 {
					if mResp.Rows[0].Columns[0].GetValueIsValid() {
						id = fmt.Sprintf("%d", mResp.Rows[0].Columns[0].GetInt32Value())
						apilog.Debugw(ctx, "insert result", "id", id)
					} else {
						apilog.Debugw(ctx, "value is not valid")
					}
				} else {
					apilog.Debugw(ctx, "empty columns")
				}
			} else {
				apilog.Debugw(ctx, "empty rows")
			}
		} else {
			apilog.Debugw(ctx, "db error", "err", mResp.Error)
		}

		// publish
		_, err = s.queueClient.Publish(ctx, &apiqueue.PublishRequest{Topic: "todo", Message: []byte("Hello microgate")})
		if err != nil {
			return apilog.ErrorWithLog(ctx, err, "failed to publish todo created")
		}
		// all ok
		return nil
	})

	return &todo.CreateTodoResponse{
		Id: id,
	}, nil
}
