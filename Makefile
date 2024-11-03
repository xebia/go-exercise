.PHONY: lint, lint-fix, test

export HOST = 127.0.0.1
export PORT = 8080

run:
	go run ./cmd/humanitechd/humanitechd.go

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix

test:
	go test -count=1 ./...