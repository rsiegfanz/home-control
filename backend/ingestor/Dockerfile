# syntax=docker/dockerfile:1

# 1. BUILD
FROM golang:1.23-alpine AS builder-backend

WORKDIR /app

ENV GO111MODULE=on

COPY ./backend/ingestor/go.mod ./backend/ingestor/go.sum ./ingestor/
COPY ./backend/sharedlib/go.mod ./backend/sharedlib/go.sum* ./sharedlib/

RUN go mod download

COPY ./backend/ingestor/cmd ./ingestor/cmd
COPY ./backend/ingestor/pkg ./ingestor/pkg

COPY ./backend/sharedlib/internal ./sharedlib/internal
COPY ./backend/sharedlib/pkg ./sharedlib/pkg

WORKDIR /app/ingestor

RUN CGO_ENABLED=0 GOOS=linux go build -o ./main ./cmd/main.go

# 2. RUNTIME
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder-backend /app/ingestor/main .

EXPOSE 8080

RUN adduser -D -g '' nonroot
USER nonroot:nonroot

# RUN
CMD ["./main", "-env=prod"]
