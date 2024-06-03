# syntax=docker/dockerfile:1

# BACKEND
FROM golang:1.22.3 as build-backend

WORKDIR app

COPY backend/go.mod backend/go.sum ./

RUN go mod download

COPY backend/cmd ./cmd
COPY backend/pkg ./pkg

RUN CGO_ENABLED=0 GOOS=linux go build -o backend-service cmd/main.go

RUN ls -l .

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage backend-service backend-service

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/docker-gs-ping"]


#FRONTEND
FROM node:20 AS frontend

WORKDIR /webapp

COPY frontend/package*.json ./package.json
COPY frontend/src ./src

RUN npm install
RUN npm run build


# RUN

CMD ["/docker-homecontrol"]
