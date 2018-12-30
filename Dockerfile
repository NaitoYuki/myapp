FROM golang:1.11.4

ENV GOPATH $GOPATH:/go/src

RUN apt-get update && \
    apt-get upgrade -y

RUN go get github.com/revel/revel && \
    go get github.com/revel/cmd/revel && \
    go get github.com/jinzhu/gorm && \
    go get github.com/go-sql-driver/mysql

RUN mkdir /go/src/revel_docker_sample

EXPOSE 9000
