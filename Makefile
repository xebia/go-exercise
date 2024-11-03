.PHONY: lint, lint-fix, test

export HOST = 127.0.0.1
export PORT = 8080
export SENDGRID_API_KEY = #


test:
	go test -count=1 ./...

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix

run:
	go run ./cmd/humanitechd/humanitechd.go
