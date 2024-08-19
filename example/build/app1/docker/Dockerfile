# syntax=docker/dockerfile:1
FROM golang:1.21-alpine3.18 AS build
WORKDIR /app
COPY ./ ./
RUN go mod download
RUN go build -o /program ./build/src/app1

#######
FROM build
WORKDIR /
COPY --from=build /program /program
ENTRYPOINT ["/program"]