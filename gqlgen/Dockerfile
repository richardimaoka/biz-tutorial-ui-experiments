# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY gqlgen.yml ./gqlgen.yml
COPY graph ./graph
COPY data ./data

RUN go build -o ./gqlgen-server

EXPOSE 8080

CMD [ "/app/gqlgen-server" ]