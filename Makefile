run: build
	@./bin/app
build:
	@go build -o bin/app cmd/wt-guided-weaponry/main.go