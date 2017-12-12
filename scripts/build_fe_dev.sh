#! /usr/bin/env bash
cd $FRONTEND_PATH
$(npm bin)/ng build --watch --output-path $GOPATH/src/$PKG_NAME/static --deploy-url /static
