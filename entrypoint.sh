#!/bin/sh

if [ "$MODE" = "production" ]; then
  /root/main
else
  go run /app/cmd/main.go
fi