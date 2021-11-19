module github.com/microgate-io/microgate-demo

go 1.16

require (
	github.com/Khan/genqlient v0.3.0
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/microgate-io/microgate-lib-go v1.0.0
	golang.org/x/net v0.0.0-20211116231205-47ca1ff31462 // indirect
	golang.org/x/sys v0.0.0-20210910150752-751e447fb3d0 // indirect
	google.golang.org/genproto v0.0.0-20200825200019-8632dd797987 // indirect
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v3 v3.0.0-20210106172901-c476de37821d // indirect
)

replace github.com/microgate-io/microgate-lib-go v1.0.0 => ../microgate-lib-go
