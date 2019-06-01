#!/usr/bin/env sh
golangci-lint run  ./... -v \
&& go test ./... -race -coverprofile=cover.out