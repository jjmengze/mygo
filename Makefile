include Makefile.common

lint: lint-tools
	golangci-lint run --config .golangci.yaml
.PHONY: lint

gen: gqlgen-tools
	# generate GQL model & resolver
	gqlgen
.PHONY: gen

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
