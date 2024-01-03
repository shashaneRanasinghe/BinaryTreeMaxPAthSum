## Build
FROM golang:1.20-alpine3.16 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build -o binaryTreeMaxPathSum ./cmd/app

## Deploy

FROM alpine:3.16.0
WORKDIR /
COPY --from=build app/binaryTreeMaxPathSum /binaryTreeMaxPathSum
ENTRYPOINT ["/binaryTreeMaxPathSum"]