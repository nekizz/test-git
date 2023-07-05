FROM golang:1.20

ADD ./bin/go-test-service /usr/bin/go-test-service

EXPOSE 1323