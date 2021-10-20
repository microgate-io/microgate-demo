run:
	go run *.go

pb:
	protoc \
	-I proto \
	--go_out=todo --go_opt=paths=source_relative \
	--go-grpc_out=todo --go-grpc_opt=paths=source_relative \
	todo.proto
	
	protoc \
	-I proto \
	--go_out=user --go_opt=paths=source_relative \
	--go-grpc_out=user --go-grpc_opt=paths=source_relative \
	user.proto