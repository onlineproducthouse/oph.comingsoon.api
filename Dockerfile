# syntax=docker/dockerfile:1

ARG IMAGE_REGISTRY_BASE_URL

FROM ${IMAGE_REGISTRY_BASE_URL}/node:22.9-alpine3.21 AS builder

LABEL maintainer="Bongani Masuku <bongani@1702tech.com>"

RUN mkdir -p /home/node/app/ && chown -R node:node /home/node/app

WORKDIR /home/node/app/

USER node

COPY --chown=node:node package.json ./

RUN npm run clean && npm i

#==================================================================

FROM builder

COPY ./src ./src
COPY ./package.json ./package.json
COPY ./tsconfig.json ./tsconfig.json

RUN npm run build

ENTRYPOINT [ "sh", "-c", "npm run serve" ]