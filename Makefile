LDFLAGS=-w
COMPILER=go build -ldflags="$(LDFLAGS)"
ENVVARS=CGO_ENABLED=0 GOOS=linux GOARCH=amd64

test:
	go test ./...

dev-postgres:
	docker run -d -e POSTGRES_PASSWORD=secret12345 -e POSTGRES_DB=vatusa -p 5432:5432 postgres:13-alpine

build:
	$(ENVVARS) $(COMPILER) -o build/api

clean:
	rm -rf build

all: test build