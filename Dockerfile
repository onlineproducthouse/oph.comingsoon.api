# syntax=docker/dockerfile:1

ARG IMAGE_REGISTRY_BASE_URL

FROM ${IMAGE_REGISTRY_BASE_URL}/node:22.9-alpine AS builder

LABEL maintainer="Bongani Masuku <bongani@1702tech.com>"

RUN mkdir -p /app && chown -R node:node /app/

COPY ./src /app/src
COPY ./yarn.lock /app/yarn.lock
COPY ./package.json /app/package.json
COPY ./tsconfig.json /app/tsconfig.json

WORKDIR /app

RUN corepack enable && yarn set version berry && yarn && yarn build

ENTRYPOINT [ "sh", "-c", "yarn serve" ]