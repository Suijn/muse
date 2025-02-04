#!/bin/bash

if [ "$1" = "server" ]; then
  echo Running server
  ./muse
elif [ "$1" = "test" ]; then
  go test ./...
fi
