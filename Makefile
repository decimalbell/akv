.PHONY: pkg cmd test cover bench clean

BIN_AKV=bin/akv

all: pkg cmd test

pkg:
	go build github.com/decimalbell/akv/cache

cmd:
	go build -o ${BIN_AKV} github.com/decimalbell/akv/cmd/akv

test:
	go test -race ./...

cover:
	go test -race -covermode=atomic -coverprofile=cover.out ./...
	go tool cover -html=cover.out

bench:
	go test -race -benchmem -bench=. github.com/decimalbell/akv/cache

clean:
	rm ${BIN_AKV}
