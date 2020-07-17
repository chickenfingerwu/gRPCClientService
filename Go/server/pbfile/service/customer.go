package service

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type Server struct {
	Db *sql.DB
}

//connect returns a SQL database connection from the pool
func (s *Server) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.Db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database->"+err.Error())
	}
	return c, nil
}

//AddCustomer insert new customer into the table
func (s *Server) AddCustomer(ctx context.Context, in *AddReq) (*Confirmation, error) {
	c, err := s.connect(ctx)
	defer c.Close()

	if err != nil {
		panic(err.Error())
	}

	query := fmt.Sprintf("INSERT INTO test_tb(id, name) VALUES('%v', '%s');", in.GetId(), in.GetName())

	_, qerr := c.ExecContext(ctx, query)

	if qerr != nil {
		panic(qerr.Error())
	}

	log.Printf("Receive message body from client: %v, %s", in.GetId(), in.GetName())
	return &Confirmation{Confirmation: "Added!"}, nil
}

//RetrieveCustomer returns a customer from from fields
func (s *Server) RetrieveCustomer(ctx context.Context, in *GetReq) (*Customer, error) {
	c, err := s.connect(ctx)
	defer c.Close()
	query := fmt.Sprintf("SELECT * FROM test_tb WHERE Id = %v;", in.GetId())
	results, err := c.QueryContext(ctx, query)
	if err != nil {
		panic(err.Error())
	}

	var customer Customer
	for results.Next() {
		err = results.Scan(customer.GetId(), customer.GetName())
		if err != nil {
			panic(err.Error())
		}
	}
	return &customer, nil
}

//DeleteCustomer deletes a customer by id
func (s *Server) DeleteCustomer(ctx context.Context, in *DeleteReq) (*Confirmation, error) {
	c, err := s.connect(ctx)
	defer c.Close()

	if err != nil {
		panic(err.Error())
	}

	query := fmt.Sprintf("DELETE FROM test_tb WHERE Id = %v;", in.GetId())
	_, qerr := c.ExecContext(ctx, query)
	if qerr != nil {
		panic(qerr.Error())
	}
	return &Confirmation{Confirmation: "Delete success!"}, nil
}

//GetAllCustomers returns all customers currently in database
func (s *Server) GetAllCustomers(ctx context.Context, in *GetAllReq) (*Customers, error) {
	c, err := s.connect(ctx)
	defer c.Close()

	if err != nil {
		panic(err.Error())
	}

	query := fmt.Sprintf("SELECT * FROM %s;", in.GetDbName())
	results, qerr := c.QueryContext(ctx, query)
	if qerr != nil {
		panic(qerr.Error())
	}
	var customers Customers
	for results.Next() {
		var cust Customer
		err = results.Scan(&cust.Id, &cust.Name)
		if err != nil {
			panic(err.Error())
		}
		customers.Customers = append(customers.Customers, &cust)
	}
	for _, value := range customers.Customers {
		fmt.Printf("Customer %v's name: %s\n", value.GetId(), value.GetName())
	}
	return &customers, nil
}
