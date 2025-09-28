#!/usr/bin/env bash

source .env && go run ./cmd/curaitor/main.go > ./backend.log 2>&1 &

(cd ./ui && pnpm install && pnpm dev) > ./frontend.log 2>&1 &
