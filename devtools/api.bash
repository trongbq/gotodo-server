#!/usr/bin/env bash

set -e

export ENV=local

go run ./cmd/api/main.go
