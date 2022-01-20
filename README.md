# VATUSA APIv2

Very much a work in progress

## Building local dev environment

- Requires docker
- Golang 1.17

### Install PostgreSQL container

```bash
make dev-postgres
```

### Run migrations and seed ratings

```bash
go run . migrate
go run . seed -f rating-seed.yaml
```

## Running tests

```bash
make test
```
