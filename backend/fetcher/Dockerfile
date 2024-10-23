# syntax=docker/dockerfile:1

# 1. BUILD
FROM golang:1.23-alpine AS builder-backend

WORKDIR /app

ENV GO111MODULE=on

COPY ./backend/fetcher/go.mod ./backend/fetcher/go.sum ./fetcher/
COPY ./backend/sharedlib/go.mod ./backend/sharedlib/go.sum* ./sharedlib/

WORKDIR /app/backend/fetcher
RUN go mod download

WORKDIR /app
COPY ./backend/fetcher/cmd ./fetcher/cmd
COPY ./backend/fetcher/pkg ./fetcher/pkg

# COPY ./backend/sharedlib/internal ./sharedlib/internal
COPY ./backend/sharedlib/pkg ./sharedlib/pkg

WORKDIR /app/fetcher

RUN CGO_ENABLED=0 GOOS=linux go build -o ./main ./cmd/main.go

# 2. RUNTIME
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder-backend /app/fetcher/main .

EXPOSE 8080

RUN adduser -D -g '' nonroot
USER nonroot:nonroot

# RUN
CMD ["./main", "-env=prod"]
