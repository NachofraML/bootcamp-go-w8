.PHONY: coverage
coverage:
	@echo "=> Running tests and generating report"
	@go test ./... -covermode=atomic -coverprofile=./coverage.out -coverpkg=./... -count=1
	@go tool cover -html=./coverage.out

.PHONY: lint
lint:
	@golangci-lint run