# syntax=docker/dockerfile:1

FROM golang:1.20-alpine as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o main /app

FROM alpine:3.11.3
COPY --from=build /app ./

ADD assets ./assets/

EXPOSE 42069

CMD ["./main"]