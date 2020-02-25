.DEFAULT_GOAL := build
NAME=arbitrage

install: # Install dependencies
	export GO111MODULE=on 
	go mod download
	go mod vendor

build:
	go build -o $(NAME) cmd/main.go 

test:
	go test ./...