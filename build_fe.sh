#! /usr/bin/env bash
cd $FRONTEND_PATH
mkdir -p /app/static
$(npm bin)/ng build --prod --build-optimizer --deploy-url /static && cp -rf $FRONTEND_PATH/dist/* /app/static
