# syntax=docker/dockerfile:1

ARG APP_NAME=ingestor

# 1. BUILD
FROM golang:1.23.1 AS builder

WORKDIR /app

ENV GO111MODULE=on

COPY ./backend/${APP_NAME}/go.mod ./backend/fetcher/go.sum ./${APP_NAME}/
COPY ./backend/${APP_NAME}/cmd ./${APP_NAME}/cmd
COPY ./backend/${APP_NAME}/pkg ./${APP_NAME}/pkg

COPY ./backend/sharedlib/go.mod ./backend/sharedlib/go.sum* ./sharedlib/
COPY ./backend/sharedlib/internal ./sharedlib/internal
COPY ./backend/sharedlib/pkg ./sharedlib/pkg

WORKDIR /app/${APP_NAME}
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./main ./cmd/main.go

# 2. RUNTIME
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/${APP_NAME}/main .

EXPOSE 8080

USER nonroot:nonroot

# RUN
CMD ["./main"]
