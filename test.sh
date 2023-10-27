#!/bin/sh

set -e

PKGS=$1

go test -timeout 30s -coverprofile=coverage.txt -covermode=atomic ${PKGS}
go tool cover -html=coverage.txt -o coverage.html
go tool cover -func coverage.txt