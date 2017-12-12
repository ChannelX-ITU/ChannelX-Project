#! /usr/bin/env bash
echo "Creating go path: ${GOPATH}/src/${PKG_NAME}"
mkdir -p $GOPATH/src/$PKG_NAME
echo "Symbolic link from: ${BACKEND_PATH} to: ${GOPATH}/src/${PKG_NAME}"
ln -s -f $BACKEND_PATH/* $GOPATH/src/$PKG_NAME/
type mysql >/dev/null 2>&1 && echo "MySQL present, setting up" && mysql_setup.sh || echo "MySQL not present, skipping setup"
echo "Doing npm install"
cd $FRONTEND_PATH
npm install
cd /app
echo "Running build_dev.sh"
build_dev.sh
