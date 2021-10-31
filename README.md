# microgate-demo

Example service using the micogate side-process.

## services

![](docs/setup.png)

## db setup

	docker run --name pgtalk-db -e POSTGRES_PASSWORD=microgate  -p 5432:5432 -d postgres
	
	psql -h localhost -p 5432 -U postgres
	
	CREATE DATABASE todo;
	
	\c todo

	CREATE TABLE tasks (
		task_id serial PRIMARY KEY,
		title VARCHAR ( 200 ) UNIQUE NOT NULL	
	);

## compile api

	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1 
	make pb

## run

### start microgate
### start user-server
### start todo-server
### run client test