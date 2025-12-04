# https://just.systems

# Run lint and test.
[parallel]
@default: lint test

# Run tests with race detector enabled.
test:
    go test -race ./...

# Run linters (staticcheck).
lint:
    go tool honnef.co/go/tools/cmd/staticcheck ./...
