#!/usr/bin/env bash

go run ./cmd/curaitor/main.go > ./backend.log 2>&1 &

(cd ./ui && pnpm dev) > ./frontend.log 2>&1 &
