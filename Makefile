# Copyright 2025 Ivan Guerreschi <ivan.guerreschi.dev@gmail.com>.
# All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY: build test run vet fmt clean

all: build test

lint:
	golangci-lint run --config .golangci.yml

run:
	go run ./...

vet:
	go vet ./...

fmt:
	gofmt -d -e -s -w .

test:
	go test -v ./...

build:
	go build -o build/gopwd main.go

clean:
	rm -rf build
