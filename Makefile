.PHONY: pkg cmd test clean

BIN_AKV=bin/akv

all: pkg cmd test

pkg:
	go build github.com/decimalbell/akv
	go build github.com/decimalbell/akv/cache
	go build github.com/decimalbell/akv/internal

cmd:
	go build -o ${BIN_AKV} github.com/decimalbell/akv/cmd/akv

test:
	go test -race -v ./...

bench:
	go test -race -v -bench=. ./...

clean:
	rm ${BIN_AKV}
