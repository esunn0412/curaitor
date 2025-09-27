#!/usr/bin/env bash

BACKEND_PID=$(pgrep -f "go run ./cmd/curaitor/main.go")

if [ -n "$BACKEND_PID" ]; then
  echo "Found Go backend process with PID: $BACKEND_PID. Terminating..."
  kill $BACKEND_PID
  sleep 1
else
  echo "Go backend process not found."
fi


FRONTEND_PID=$(pgrep -f "pnpm dev")

if [ -n "$FRONTEND_PID" ]; then
  echo "Found pnpm dev process with PID: $FRONTEND_PID. Terminating..."
  kill $FRONTEND_PID
  sleep 1
else
  echo "pnpm dev process not found."
fi

echo "frontend and backend killed"
