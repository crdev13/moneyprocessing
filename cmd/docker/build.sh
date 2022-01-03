#! /bin/bash

rm service

env GOOS=linux GOARCH=386 go build .

mv docker service