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

	// write to database and return new id

	// begin
	bResp, err := s.dbClient.Begin(ctx, new(apidb.BeginRequest))
	if err != nil {
		return nil, apilog.ErrorWithLog(ctx, err, "failed to begin db transaction", "db", "begin")
	}

	// mutate
	mResp, err := s.dbClient.Mutate(ctx, &apidb.MutationRequest{
		TransactionId: bResp.TransactionId,
		Sql:           fmt.Sprintf("INSERT INTO public.tasks (title) VALUES ('hello microgate %v') RETURNING task_id", time.Now())})
	if err != nil {
		apilog.Debugw(ctx, "Rollback")
		_, _ = s.dbClient.Rollback(ctx, &apidb.RollbackRequest{TransactionId: bResp.TransactionId})
		return nil, apilog.ErrorWithLog(ctx, err, "failed to execute mutation", "db", "mutate")
	}

	// TEMP hack
	apilog.Debugw(ctx, "insert result", "rows", mResp.Rows)
	var id string
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

	// commit
	_, err = s.dbClient.Commit(ctx, &apidb.CommitRequest{TransactionId: bResp.TransactionId})
	if err != nil {
		return nil, apilog.ErrorWithLog(ctx, err, "failed to commit db transaction", "db", "commit")
	}

	// publish
	_, err = s.queueClient.Publish(ctx, &apiqueue.PublishRequest{Topic: "todo", Message: []byte("Hello microgate")})
	if err != nil {
		return nil, apilog.ErrorWithLog(ctx, err, "failed to publish todo created")
	}

	return &todo.CreateTodoResponse{
		Id: id,
	}, nil
}
