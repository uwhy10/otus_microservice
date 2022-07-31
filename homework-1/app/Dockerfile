# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY go.mod  ./
COPY main.go ./
RUN go mod download
RUN go build -o homework1


FROM alpine:3.15
WORKDIR /app/
COPY --from=builder /app/homework1 ./
CMD ["/app/homework1"]