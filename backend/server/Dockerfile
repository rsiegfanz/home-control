# syntax=docker/dockerfile:1

# 1. BUILD BACKEND
FROM golang:1.23.1 AS builder-backend

WORKDIR /app

ENV GO111MODULE=on

COPY ./backend/server/go.mod ./backend/server/go.sum ./server/
COPY ./backend/server/cmd ./server/cmd
COPY ./backend/server/pkg ./server/pkg

COPY ./backend/sharedlib/go.mod ./backend/sharedlib/go.sum* ./sharedlib/
COPY ./backend/sharedlib/internal ./sharedlib/internal
COPY ./backend/sharedlib/pkg ./sharedlib/pkg

WORKDIR /app/server
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./main ./cmd/main.go

# 2. BUILD FRONTEND
FROM node:20 AS builder-frontend

WORKDIR /app

COPY frontend/ ./

RUN npm install
RUN npm run build:prod

# 3. RUNTIME
FROM alpine:latest AS runtime

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder-backend /app/server/main .
COPY --from=builder-frontend /app/dist/home-control/browser ./webapp
# RUN mkdir -p /data

EXPOSE 8080

# USER nonroot:nonroot

# RUN
CMD ["./main", "-env=prod"]

############################
# FRONTEND
############################

# FROM node:20 AS builder-frontend

# WORKDIR /app

# COPY frontend/ ./

# RUN npm install
# RUN npm run build:prod

# # COPY
# FROM gcr.io/distroless/base-debian11 AS build-release-stage

# WORKDIR /app
# COPY --from=builder-backend /backend-service ./backend-service
# COPY --from=builder-backend /data ./data

# COPY --from=builder-frontend app/dist/home-control/browser ./webapp


# EXPOSE 8080

# USER nonroot:nonroot


# RUN
# CMD ["/app/backend-service"]
