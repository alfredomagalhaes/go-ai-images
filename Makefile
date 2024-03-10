run: build
	@./bin/imagegenerator

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download
	@npm install -D tailwindcss
	@npm install -D daisyui@latest

build:
	npx tailwindcss -i view/css/app.css -o public/styles.css
	@templ generate view
	@go build -o bin/imagegenerator main.go

up: ## Databae migration up
	@go run cmd/migrate/main.go up

drop: 
	@go run cmd/drop/main.go up
