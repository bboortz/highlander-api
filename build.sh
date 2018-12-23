#!/bin/bash

set -e
set -u


if [ -n "${DEV:-}" ]; then
    go build -o highlander cmd/main.go
else
    docker build -t highlander -f build/Dockerfile .
fi
