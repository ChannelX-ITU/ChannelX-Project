#! /usr/bin/env bash
cd $GOPATH/src/$PKG_NAME && go get && go build && cp -f $GOPATH/src/$PKG_NAME/$(basename $PKG_NAME) /app/server
