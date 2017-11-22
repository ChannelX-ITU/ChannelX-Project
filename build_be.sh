#! /usr/bin/env bash
cd $BACKEND_PATH
mkdir -p $GOPATH/src/$PKG_NAME
cp -rf . $GOPATH/src/$PKG_NAME
cd $GOPATH/src/$PKG_NAME && go get && go build && cp -f $GOPATH/src/$PKG_NAME/$(basename $PKG_NAME) /app/server
