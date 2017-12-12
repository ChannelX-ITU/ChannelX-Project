#! /usr/bin/env bash
cd $FRONTEND_PATH
mkdir -p /app/static
$(npm bin)/ng build --prod --build-optimizer --output-path /app/static --deploy-url /static
