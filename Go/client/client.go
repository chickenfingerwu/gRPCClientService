package main

import (
	"./pbfile/service"
	"context"
	"encoding/json"
	"fmt"
	//"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

//port is for server, frontendPort is for frontend client
const (
	port         = "server:8080"
	frontendPort = ":9090"
)

//homePage function
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

//Handler stores connection and the service client to the server,
//also handles any requests made to the api
type Handler struct {
	conn *grpc.ClientConn
	c    service.ServerServiceClient
}

//handleRequests handles all request made to api
func (h *Handler) handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/customers", h.returnAllCustomers)
	log.Fatal(http.ListenAndServe(frontendPort, nil))
}

//returnAllCustomers all customers from database
func (h *Handler) returnAllCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllCustomers")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	response, err := h.c.GetAllCustomers(context.Background(), &service.GetAllReq{
		DbName: "test_tb",
	})
	if err != nil {
		log.Fatalf("Error when calling GetAllCustomers: %s", err)
	}
	log.Printf("Response from server: %s", response.Customer)
	json.NewEncoder(w).Encode(response.Customer)

	//out, err := proto.Marshal(response)
	//if err != nil {
	//	log.Fatalln("Failed to encode response from server: ", err)
	//}
	//
	//
	//_, werr := w.Write(out)
	//if werr != nil {
	//	log.Fatalln("Writing to client failed: ", werr)
	//}
}

func main() {
	var conn *grpc.ClientConn

	//dial the server
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	//keep connection open until program exit
	defer conn.Close()

	c := service.NewServerServiceClient(conn)

	//initialize the handler
	handler := Handler{conn, c}
	handler.handleRequests()
}
