LDFLAGS=-w
COMPILER=go build -ldflags="$(LDFLAGS)"
ENVVARS=CGO_ENABLED=0 GOOS=linux GOARCH=amd64

test:
	go test ./...

dev-containers:
	docker run -d -e POSTGRES_PASSWORD=secret12345 -e POSTGRES_DB=vatusa -p 5432:5432 postgres:13-alpine
	docker run -d -p 6379:6379 redis:6-alpine

dev-jwks:
	bash scripts/generate-jwks.sh

build:
	$(ENVVARS) $(COMPILER) -o build/api

clean:
	rm -rf build

dev-postgres: dev-containers
all: test build