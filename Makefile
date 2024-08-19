run: build
	@./bin/app

build:
	@go build -o bin/app cmd/wt-guided-weaponry/main.go

css:
	npx tailwindcss -i views/css/app.css -o public/styles.css --watch