#! /usr/bin/env bash
cd $GOPATH/src/$PKG_NAME && go get -v && go build -v -o /app/server
