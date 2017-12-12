FROM node:8-alpine as builder

WORKDIR /app

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh musl-dev go

RUN apk add --no-cache ca-certificates

RUN mkdir -p /app/backend /app/frontend

RUN mkdir -p /app/go 
ENV GOPATH /app/go 
RUN mkdir -p $GOPATH/pkg $GOPATH/bin $GOPATH/src

ENV FRONTEND_PATH /app/frontend
ENV BACKEND_PATH /app/backend
ENV PATH="/app:${GOPATH}/bin:${PATH}"

RUN go get -v github.com/achiku/wbs

COPY my.cnf /etc/mysql/my.cnf
COPY scripts/mysql_setup.sh /app/mysql_setup.sh
COPY scripts/startup.sh /app/startup.sh
COPY scripts/build_be.sh /app/build_backend.sh
COPY scripts/build_fe.sh /app/build_frontend.sh
COPY scripts/build_fe_dev.sh /app/build_frontend_dev.sh
COPY scripts/build_be_dev.sh /app/build_backend_dev.sh
COPY scripts/build_dev.sh /app/build_dev.sh
COPY scripts/build.sh /app/build.sh
COPY scripts/wait-for-it.sh /app/wait-for-it.sh
RUN chmod +x startup.sh mysql_setup.sh build_backend.sh build_frontend.sh build_frontend_dev.sh build_backend_dev.sh build_dev.sh build.sh wait-for-it.sh

EXPOSE 4200
EXPOSE 6969

ENTRYPOINT /bin/bash

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


