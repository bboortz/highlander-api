#!/bin/bash

set -e
set -u


if [ -n "${DEV:-}" ]; then
    go run cmd/main.go
else
    docker run  -i -t --rm --name highlander  -p 8444:8444 highlander
fi