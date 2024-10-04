generate-proto: 
	buf generate
	@echo \# source code is generated

test:
	go test ./...  -count=1 -v -cover -race

lint: ## Run linter
	golangci-lint run ./...