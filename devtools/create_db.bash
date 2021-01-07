#!/usr/bin/env bash

set -e

mysql -u root -h 127.0.0.1 -e "CREATE DATABASE IF NOT EXISTS gotodo CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;"

