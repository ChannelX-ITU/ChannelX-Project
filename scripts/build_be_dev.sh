#! /usr/bin/env bash
cd $GOPATH/src/$PKG_NAME && go get -v && wbs -c wbs.config.toml
