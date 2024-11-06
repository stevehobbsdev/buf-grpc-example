.PHONY: generate clean build run

build:
	@go build -o out/main main.go

run:
	@go run main.go

generate:
	@buf generate

clean:
	@rm -rf gen