FROM node:8-alpine as builder

WORKDIR /app

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh musl-dev go mysql mysql-client

RUN mkdir -p /app/backend /app/frontend

ENV FRONTEND_PATH /app/frontend
ENV BACKEND_PATH /app/backend
ENV PATH="/app:${PATH}"

RUN mkdir -p /app/go 
ENV GOPATH /app/go 
RUN mkdir -p $GOPATH/pkg $GOPATH/bin $GOPATH/src

COPY my.cnf /etc/mysql/my.cnf
COPY mysql_setup.sh /app/mysql_setup.sh
COPY startup.sh /app/startup.sh
COPY build_be.sh /app/build_backend.sh
COPY build_fe.sh /app/build_frontend.sh
COPY build_fe_dev.sh /app/build_frontend_dev.sh
RUN chmod u+x startup.sh mysql_setup.sh build_backend.sh build_frontend.sh build_frontend_dev.sh

EXPOSE 4200
EXPOSE 6969

ENTRYPOINT startup.sh && /bin/bash

# FROM alpine
# 
# RUN apk add --no-cache bash
# 
# WORKDIR /app
# 
# COPY --from=builder /app/backend /app/backend
# COPY --from=builder /app/static/ /app/static/*
# 
# COPY wait-for-it.sh .
# RUN chmod u+x wait-for-it.sh
# 
# CMD ["./wait-for-it.sh", "mysql:3306", "--timeout=0", "--", "/bin/sh"]


