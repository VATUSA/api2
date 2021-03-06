# VATUSA APIv2

Very much a work in progress

## Building local dev environment

### Requirements

- Requires docker
- Golang 1.17
- A Linux environment (WSL, Linux, OS X [should work?])

### Install and run

1. Install PostgreSQL and Redis containers

    ```bash
    make dev-containers
    ```

    **Note** In the development environment, Redis will not use a password. Either leave this blank in the YAML or comment out the password line.

2. Run migrations and seed ratings (more soon)

    ```bash
    go run . migrate
    go run . seed -f rating-seed.yaml
    ```

3. Create jwks (dev environment only!)

    ```bash
    make dev-jwks
    ```

4. Create config and edit

    ```bash
    cp config.yaml.example config.yaml
    ```

5. Generate docs (optional if building binary using `make build`, otherwise required)

   ```bash
   make build-docs
   ```

6. Generate binary (optional if using `go run .`)

    ```bash
    make build
    ```

7. Starting the API

    ```bash
    go run . server
    ## OR ##
    ./build/api server
    ```

## Running tests

```bash
make test
```

## Structure

- cmd/cmd.go - Common functions used by multiple subcommands
- cmd/(subdir) - Subcommands for the api binary
- docs/ - Swaggo generated files [best not to touch]
- internal/ - Packages that are only usable by the api binary, they have dependencies that are unlikely to be addressed by other projects or are project specific
- pkg/ - Packages that should have minimal requirements to maintain portability, these can be used in other VATUSA projects or by external groups
- scripts/ - These are helper scripts that are used external to the API binary (development env building, container image building, etc.)
- static/ - Static html files

## Contributors

- [Daniel Hawton](https://github.com/dhawton)

## License

[BSD 3-Clause](LICENSE)