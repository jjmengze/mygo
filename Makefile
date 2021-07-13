lint:
	@command -v golangci-lint > /dev/null 2>&1 || (cd $${TMPDIR} && go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.38.0)
	golangci-lint run --config .golangci.yaml
.PHONY: lint


test:
	@go test \
		-count=1 \
		-short \
		-timeout=5m \
		./...
.PHONY: test


test-coverage:
	@go tool cover -func=./coverage.out
.PHONY: test-coverage