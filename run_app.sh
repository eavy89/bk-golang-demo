#!/bin/bash

go mod tidy
CGO_ENABLED=1 go build -o out/myapp ./src/app

./out/myapp