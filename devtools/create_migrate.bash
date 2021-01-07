#!/usr/bin/env bash

set -e

migrate create -ext sql -dir migrations -seq $1
