package service

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
)

type Server struct {
	Db *sql.DB
}

type InvalidTypeError struct {
	What string
}

type QueryError struct {
	Err  error
	What string
	Code int
}

type appHandler func(http.ResponseWriter, *http.Request) error

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		if e, ok := err.(QueryError); ok {
			http.Error(w, e.What, e.Code)
			log.Fatalf("%v", e.Error())
		}
	}
}

func (e QueryError) Error() string {
	return fmt.Sprintf("Query error: %s", e.What)
}

func (e InvalidTypeError) Error() string {
	return fmt.Sprintf("Invalid type: %s", e.What)
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
	t := fmt.Sprintf("%T", in.Name)
	if t != "string" {
		return nil, &InvalidTypeError{
			What: fmt.Sprintf("Expected string, got %s.", t),
		}
	}

	c, err := s.connect(ctx)
	defer c.Close()

	if err != nil {
		panic(err.Error())
	}

	result, qerr := c.ExecContext(ctx, "INSERT INTO Customer(Name) VALUES( ? )", in.GetName())
	if qerr != nil {
		panic(qerr.Error())
	}

	lastInID, err := result.LastInsertId()
	if err != nil {
		panic(qerr.Error())
	}

	return &Confirmation{LastInsertID: lastInID, Confirmation: "Added!"}, nil
}

//RetrieveCustomer returns a customer from from fields
func (s *Server) RetrieveCustomer(ctx context.Context, in *GetReq) (*Customer, error) {
	t := fmt.Sprintf("%T", in.Id)
	if t != "int64" && t != "int32" && t != "int" {
		return nil, &InvalidTypeError{
			What: fmt.Sprintf("Expected int, int32 or int64, got %s.", t),
		}
	}

	c, err := s.connect(ctx)
	defer c.Close()

	results, err := c.QueryContext(ctx, "SELECT * FROM Customer WHERE CustomerID = ?", in.GetId())
	defer results.Close()

	if err != nil {
		log.Fatalf("Encountered error: %v", err.Error())
	}

	var customer Customer

	// check if result returned is empty or not
	if ok := results.Next(); !ok {
		return nil, &QueryError{
			Err:  err,
			What: "Result returned empty!",
			Code: 404,
		}
	} else {
		err = results.Scan(&customer.Id, &customer.Name)
		if err != nil {
			panic(err.Error())
		}
	}

	// if rows not empty then read rows
	for results.Next() {
		err = results.Scan(&customer.Id, &customer.Name)
		if err != nil {
			panic(err.Error())
		}
	}

	return &customer, nil
}

//DeleteCustomer deletes a customer by id
func (s *Server) DeleteCustomer(ctx context.Context, in *DeleteReq) (*Confirmation, error) {
	t := fmt.Sprintf("%T", in.Id)
	if t != "int64" && t != "int32" && t != "int" {
		return nil, &InvalidTypeError{
			What: fmt.Sprintf("Expected int, int32 or int64, got %s.", t),
		}
	}

	c, err := s.connect(ctx)
	defer c.Close()

	if err != nil {
		panic(err.Error())
	}

	_, qerr := c.ExecContext(ctx, "DELETE FROM Customer WHERE CustomerID = ?", in.GetId())
	if qerr != nil {
		return nil, QueryError{
			Err:  err,
			What: "Customer doesn't exist or has already been deleted!",
			Code: 404,
		}
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

	query := fmt.Sprintf("SELECT * FROM Customer;")
	results, qerr := c.QueryContext(ctx, query)
	defer results.Close()

	if qerr != nil {
		return nil, &QueryError{
			Err:  err,
			What: "There are no customers!",
			Code: 404,
		}
	}

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
	return &customers, nil
}
