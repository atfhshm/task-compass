.DEFAULT_GOAL := run

test:
	go test ./... -v
test-cover:
	go test ./... -coverprofile=coverage
cover-report:
	go tool cover -html=coverage

run:
	air

.PHONY: test test-cover cover-report