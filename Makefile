test:
	go test ./...

dev-postgres:
	docker run -d -e POSTGRES_PASSWORD=secret12345 -e POSTGRES_DB=vatusa -p 5432:5432 postgres:13-alpine

all: test