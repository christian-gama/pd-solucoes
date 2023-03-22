FROM golang:1.20-alpine3.17 AS base

FROM base AS api
ARG WORKDIR
WORKDIR $WORKDIR
COPY . ./
