package main

import (
	"context"

	apilog "github.com/microgate-io/microgate-lib-go/v1/log"
	apiqueue "github.com/microgate-io/microgate-lib-go/v1/queue"
)

type MessageHandlingServiceImpl struct {
	apiqueue.UnimplementedMessageHandlingServiceServer
}

func (s *MessageHandlingServiceImpl) HandleMessage(ctx context.Context, req *apiqueue.HandleMessageRequest) (*apiqueue.HandleMessageResponse, error) {
	apilog.Debugw(ctx, "HandleMessage", "req", req)
	return nil, nil
}
