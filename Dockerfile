# syntax=docker/dockerfile:1

ARG IMAGE_REGISTRY_BASE_URL

FROM ${IMAGE_REGISTRY_BASE_URL}/golang:1.19-alpine AS builder

LABEL maintainer="Bongani Masuku <bongani@1702tech.com>"

RUN mkdir -p /app

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN go install github.com/swaggo/swag/cmd/swag

#==================================================================

FROM builder

ARG TARGETOS
ARG TARGETARCH

ENV OPH_BIN_FOLDER=./bin/${TARGETOS}/${TARGETARCH}

COPY ./oph.api.comingsoon.app ./oph.api.comingsoon.app
COPY ./oph.api.comingsoon.infrastructure ./oph.api.comingsoon.infrastructure
COPY ./oph.api.comingsoon.ioc ./oph.api.comingsoon.ioc
COPY ./oph.api.comingsoon.service ./oph.api.comingsoon.service
COPY ./oph.api.comingsoon.testhelpers ./oph.api.comingsoon.testhelpers
COPY ./oph.api.comingsoon.util ./oph.api.comingsoon.util

RUN $(go env GOPATH)/bin/swag init -g oph.api.comingsoon.app/app/app.go
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-s -w" -o ${OPH_BIN_FOLDER}/app ./oph.api.comingsoon.app/main.go

ENTRYPOINT [ "sh", "-c", "$(echo ${OPH_BIN_FOLDER}/app)" ]