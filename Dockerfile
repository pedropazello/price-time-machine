FROM golang:1.17.2

RUN go mod init

RUN mkdir /app
WORKDIR /app

COPY . /app
