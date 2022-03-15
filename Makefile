MG_USER="admin"
MG_PASS="admin"
MG_ADDR="localhost"
MG_PORT="27017"

go-build:
	go build -o bin/main main.go

go-run:
	go run main.go

docker-stop:
	docker stop $$(docker ps -a -q)
	docker rm smart_triage_mongo

docker-mongo:
	docker run --name smart_triage_mongo -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=${MG_USER} -e MONGO_INITDB_ROOT_PASSWORD=${MG_PASS} -d mongo:latest

docker-build:
	docker build -t core-app .

docker-exec:
	docker run --network host -p 8080 -e MONGO_ADDR=${MG_ADDR} -e MONGO_PORT=${MG_PORT} -e MONGO_USER=${MG_USER} -e MONGO_PASS=${MG_PASS}  core-app:latest

go-exec:
	./bin/main

compile:
	GOOS=windows go build -o bin/main-windows main.go
	GOOS=linux go build -o bin/main-linux main.go

all: go-build go-exec

all-docker: docker-stop docker-mongo docker-build docker-exec