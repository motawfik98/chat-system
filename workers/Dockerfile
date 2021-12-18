# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY . ./

RUN go mod download

WORKDIR workers

RUN go build -o /docker-workers

CMD ["/docker-workers"]