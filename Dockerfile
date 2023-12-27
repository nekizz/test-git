FROM centos:7

ADD ./bin/go-test-service /usr/bin/go-test-service

EXPOSE 1323