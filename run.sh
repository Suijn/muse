#!/bin/bash

if [ "$1" = "api" ]; then
  echo Running DB migrations
  dbmate up
  echo Running api
  ./api
elif [ "$1" = "test" ]; then
  echo Running DB migrations
  dbmate up
  go test github.com/muse/tests/common
  go test github.com/muse/tests/users
fi
