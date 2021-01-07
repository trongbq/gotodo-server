#!/usr/bin/env bash

set -e

migrate -path migrations/ -database "mysql://root@tcp(127.0.0.1:3306)/gotodo" $@
