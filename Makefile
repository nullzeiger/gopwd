# Copyright 2022 Ivan Guerreschi <ivan.guerreschi.dev@gmail.com>.
# All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

BINARY_NAME=gopwd
BINARY_PATH=build

all: fmt  vet test lint

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
	go build -o ${BINARY_PATH}/${BINARY_NAME} cmd/gopwd/main.go

clean:
	go clean
	rm -r ${BINARY_PATH}/${BINARY_NAME}
