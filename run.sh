#!/bin/bash

if [ "$1" = "server" ]; then
  echo Running DB migrations
  dbmate up
  echo Running server
  ./muse
elif [ "$1" = "test" ]; then
  echo Running DB migrations
  dbmate up
  go test github.com/muse/tests/common
  go test github.com/muse/tests/users
fi
