#!/bin/bash

if [ "$1" = "server" ]; then
  echo Running server
  ./muse
elif [ "$1" = "test" ]; then
  go test github.com/muse/tests/common
  go test github.com/muse/tests/users
fi
