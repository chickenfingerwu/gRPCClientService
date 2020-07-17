package main

import (
	"./pbfile/service"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":8080"
)

func main() {
	fmt.Println("Running!")

	//listen for incoming message
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//open connection to database
	db, err := sql.Open("mysql", "root:root@tcp(db:3306)/test_db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//initialize Server to handle queries to db,
	//initialize gRPC server
	s := service.Server{Db: db}
	grpcServer := grpc.NewServer()

	service.RegisterServerServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
