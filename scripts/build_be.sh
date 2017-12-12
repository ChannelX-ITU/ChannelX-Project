#! /usr/bin/env bash
cd $GOPATH/src/$PKG_NAME && go get && go build -o /app/server
server
