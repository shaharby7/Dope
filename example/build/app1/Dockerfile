# syntax=docker/dockerfile:1
FROM golang:1.21-alpine3.18 AS build
WORKDIR /workdir
COPY ./ ./
RUN echo $(ls -1 symlinks/Dope)
RUN go mod tidy
RUN go build -o /program ./build/app1/src

#######
FROM build
WORKDIR /
COPY --from=build /program /program
ENTRYPOINT ["/program"]