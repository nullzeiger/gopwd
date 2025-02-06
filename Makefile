# Copyright 2025 Ivan Guerreschi <ivan.guerreschi.dev@gmail.com>.
# All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY: lint
lint:
	golangci-lint run --config .golangci.yml

.PHONY: run
run:
	go run ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	gofmt -d -e -s -w .

.PHONY: test
test:
	go test -v ./...

.PHONY: build
build:
	go build -o build/gopwd main.go

.PHONY: clean
clean:
	rm -rf build
