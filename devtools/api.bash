#!/usr/bin/env bash

set -e

export ENV=local
export DATABASE_URI="root:@tcp(127.0.0.1:3306)/gotodo?parseTime=true"

go run ./cmd/api/main.go
