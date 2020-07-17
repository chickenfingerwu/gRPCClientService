FROM golang:latest

RUN apt-get update
RUN apt-get install vim -y
RUN go get -u "github.com/golang/protobuf/protoc-gen-go"
RUN go get -u "github.com/go-sql-driver/mysql"
RUN go get "google.golang.org/grpc"
RUN go get "golang.org/x/net/context"