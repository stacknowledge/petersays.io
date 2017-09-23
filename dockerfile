FROM golang:1.9.0
MAINTAINER Pedro Fernandes <stacknowledge@gmail.com>

ADD . /go/src/github.com/stacknowledge/petersays.io
WORKDIR /go/src/github.com/stacknowledge/petersays.io

RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -u github.com/Unknwon/bra
RUN dep ensure
RUN go install

EXPOSE 8080 8080

ENTRYPOINT ["bra",  "run"]

