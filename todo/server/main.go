package main

import (
	"context"
	"net"

	"github.com/microgate-io/microgate-demo/todo"
	"github.com/microgate-io/microgate-demo/user"
	v1 "github.com/microgate-io/microgate-lib-go/v1"
	apiconfig "github.com/microgate-io/microgate-lib-go/v1/config"
	apidb "github.com/microgate-io/microgate-lib-go/v1/db"
	apilog "github.com/microgate-io/microgate-lib-go/v1/log"
	apiqueue "github.com/microgate-io/microgate-lib-go/v1/queue"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// get connection
	conn := v1.DialMicrogate()

	// setup logging
	apilog.InitClient(conn)

	// get configuration
	config := apiconfig.GetConfig(conn)

	// not sure if this is right location
	apilog.GlobalDebug = config.FindBool("global_debug")

	grpcServer := grpc.NewServer()

	// for processing incoming sync requests
	registerTodo(grpcServer, conn, config)

	// for processing incoming async messages
	registerAsyncHandler(grpcServer, conn, config)

	// for processing incoming sync HTTP transcoded requests
	registerTranscodingHandler(grpcServer, conn, config)

	// start gRPC server
	addr := ":9090"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		apilog.Infow(context.Background(), "failed to listen", "err", err)
	}
	apilog.Infow(context.Background(), "serving gRPC", "addr", addr)

	// expose filedescriptors of all services
	reflection.Register(grpcServer)

	grpcServer.Serve(lis)
}

// create a new Todo service and register it to the server.
func registerTodo(s *grpc.Server, conn *grpc.ClientConn, config *apiconfig.Configuration) {
	todoService := &TodoImpl{
		dbClient:    apidb.NewDatabaseServiceClient(conn),
		userClient:  user.NewUserServiceClient(conn),
		queueClient: apiqueue.NewQueueingServiceClient(conn),
	}
	todo.RegisterTodoServiceServer(s, todoService)
}

// create a new PubSub message handler service and register it to the server.
func registerAsyncHandler(s *grpc.Server, conn *grpc.ClientConn, config *apiconfig.Configuration) {
	messageService := &MessageHandlingServiceImpl{}
	apiqueue.RegisterMessageHandlingServiceServer(s, messageService)

	// subscribe to receive messages on the handler
	c := apiqueue.NewQueueingServiceClient(conn)
	// this could cause an immediate callback on the message handling service
	_, err := c.Subscribe(context.Background(), &apiqueue.SubscribeRequest{SubscriptionId: "my-subscription-from-config"})
	if err != nil {
		apilog.Errorw(context.Background(), "failed to subscribe", "err", err)
	}
}

func registerTranscodingHandler(s *grpc.Server, conn *grpc.ClientConn, config *apiconfig.Configuration) {
	// apihttpjson.Register...Server(s)
}
