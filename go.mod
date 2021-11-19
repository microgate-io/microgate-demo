module github.com/microgate-io/microgate-demo

go 1.16

require (
	github.com/microgate-io/microgate-lib-go v1.0.0
	github.com/shurcooL/graphql v0.0.0-20200928012149-18c5c3165e3a // indirect
	golang.org/x/net v0.0.0-20211116231205-47ca1ff31462 // indirect
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
)

replace github.com/microgate-io/microgate-lib-go v1.0.0 => ../microgate-lib-go
