#!/usr/bin/env bash

BACKEND_PID=$(lsof -ti :9000)

if [ -n "$BACKEND_PID" ]; then
  echo "Found Go backend process with PID: $BACKEND_PID. Terminating..."
  kill -9 $BACKEND_PID
  sleep 1
else
  echo "Go backend process not found."
fi


FRONTEND_PID=$(lsof -ti :3000)

if [ -n "$FRONTEND_PID" ]; then
  echo "Found pnpm dev process with PID: $FRONTEND_PID. Terminating..."
  kill -9 $FRONTEND_PID
  sleep 1
else
  echo "pnpm dev process not found."
fi

echo "frontend and backend killed"
