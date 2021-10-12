FROM golang:1.17.2

RUN mkdir /app
WORKDIR /app

COPY . /app
