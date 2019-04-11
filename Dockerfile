FROM golang:1.11-alpine AS build_base
RUN apk add bash ca-certificates git gcc g++ libc-dev

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go get github.com/pilu/fresh

COPY . .
EXPOSE 4000
CMD cd server; fresh server.go
