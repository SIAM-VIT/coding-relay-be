run:
	go run cmd/main.go

postgres:
	docker run --name postgres12 -p 5433:5432 -e POSTGRES_PASSWORD=abc -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=postgres mpl-db

dropdb:
	docker exec -it postgres12 dropdb 

dockerfiledb:
	docker exec -it backend-db-1 bash

tagdocker:
	docker tag coding-relay-be guptaakshat/coding-relay-be:latest

builddocker:
	docker build -t guptaakshat/coding-relay-be:latest .

rundocker:
	docker run -p 8080:8080 guptaakshat/coding-relay-be:latest

deploydocker:
	docker push guptaakshat/coding-relay-be:latest

redisterminal:
	docker exec -it some-redis redis-cli

createredis:
	docker run -d --name some-redis -p 6379:6379 redis redis-server --requirepass abc

.PHONY: run postgres createdb dropdb dockerfiledb builddocker rundocker deploydocker tagdocker redisterminal createredis