module github.com/microgate-io/microgate-demo

go 1.16

require (
	github.com/microgate-io/microgate-lib-go v1.0.0
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
)

replace github.com/microgate-io/microgate-lib-go v1.0.0 => ../microgate-lib-go
