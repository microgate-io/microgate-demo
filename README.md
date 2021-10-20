# microgate-demo

Example service using the micogate side-process.

## db setup

docker run --name pgtalk-db -e POSTGRES_PASSWORD=microgate  -p 5432:5432 -d postgres
psql -h localhost -p 5432 -U postgres
CREATE DATABASE todo;
CREATE TABLE tasks (
	task_id serial PRIMARY KEY,
	title VARCHAR ( 50 ) UNIQUE NOT NULL	
);