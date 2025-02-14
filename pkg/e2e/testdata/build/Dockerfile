# syntax=docker/dockerfile:1
FROM golang:1.22-alpine3.18 AS build

ARG APP_NAME

WORKDIR /app
COPY ./ ./
RUN go mod download
RUN go build -o /program ./build/src/${APP_NAME}

#######
FROM build
WORKDIR /
COPY --from=build /program /program
ENTRYPOINT ["/program"]