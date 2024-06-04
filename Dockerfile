# syntax=docker/dockerfile:1

# BACKEND
FROM golang:1.22.3 as build-backend

WORKDIR app

COPY backend/go.mod backend/go.sum ./

RUN go mod download

COPY backend/cmd ./cmd
COPY backend/pkg ./pkg

RUN mkdir -p /data

RUN CGO_ENABLED=0 GOOS=linux go build -o /backend-service cmd/main.go


#FRONTEND
FROM node:20 AS build-frontend

WORKDIR app

COPY frontend/ ./

RUN npm install
RUN npm run build:prod

# COPY
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR app
COPY --from=build-backend /backend-service ./backend-service
COPY --from=build-backend /data ./data
COPY /backend/config.yaml ./config.yaml

COPY --from=build-frontend app/dist/home-control/browser ./webapp


EXPOSE 8080

USER nonroot:nonroot


# RUN
CMD ["/app/backend-service"]
