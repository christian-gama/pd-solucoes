FROM golang:1.20 AS base

FROM base AS api
RUN go install github.com/codegangsta/gin@latest
ARG WORKDIR
WORKDIR $WORKDIR
COPY . ./

FROM base AS test
ARG WORKDIR
WORKDIR $WORKDIR
COPY . ./