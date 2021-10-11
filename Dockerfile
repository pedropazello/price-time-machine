FROM golang:1.17.2

RUN go mod init
RUN go get -u -v \
        github.com/antchfx/htmlquery

RUN mkdir /app
WORKDIR /app

COPY . /app
