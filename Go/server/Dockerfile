FROM golang:latest

RUN apt-get update
RUN apt-get install vim -y
#RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.12.1/protoc-3.12.1-linux-x86_64.zip
#RUN unzip protoc-3.12.1-linux-x86_64.zip -d $HOME/.local
#RUN export PATH="$PATH:$HOME/.local/bin"
RUN go get -u "github.com/golang/protobuf/protoc-gen-go"
RUN go get -u "github.com/go-sql-driver/mysql"
RUN go get -u "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
RUN go get "google.golang.org/grpc"
RUN go get "golang.org/x/net/context"
