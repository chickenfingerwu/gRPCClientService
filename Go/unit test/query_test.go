package main

import (
	"../server/pbfile/service"
	"context"
	"fmt"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"log"
	"testing"
)

const (
	port = ":8080"
)

// initConn initialize connection to server, runs at beginning of each test function,
// returns a closure so that we can defer the closing of
// connection until after each function finish its test
func initConn() func() (service.Server, sqlmock.Sqlmock) {
	var s service.Server
	var mock sqlmock.Sqlmock

	return func() (service.Server, sqlmock.Sqlmock) {
		db, mock1, err := sqlmock.New()
		mock = mock1
		if err != nil {
			log.Fatalf("did not connect to mock server: %s", err)
		}

		s = service.Server{Db: db}

		return s, mock
	}
}

// TestGetAllCustomers tests logic of the GetAllCustomers function
func TestGetAllCustomers(t *testing.T) {

	// initialize connection
	initSet := initConn()
	s, mock := initSet()
	defer s.Db.Close()

	rows := sqlmock.NewRows([]string{"CustomerID", "Name"}).
		AddRow(1, "Son").
		AddRow(2, "Dat")

	mock.ExpectQuery("SELECT *").WillReturnRows(rows)

	// get all customers
	allCustomers, err := s.GetAllCustomers(context.Background(), &service.GetAllReq{
		DbName: "Customer",
	})

	if err != nil {
		t.Errorf("query failed: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	for _, value := range allCustomers.Customers {
		fmt.Printf("id: %d, name: %v\n", value.Id, value.Name)
	}
}

// TestAddCustomers tests logic of the AddCustomer function
func TestAddCustomer(t *testing.T) {

	// initialize connection
	initSet := initConn()
	s, mock := initSet()
	defer s.Db.Close()

	sqlmock.NewRows([]string{"CustomerID", "Name"}).
		AddRow(1, "Son").
		AddRow(2, "Dat")

	mock.ExpectExec("INSERT INTO ").WithArgs("Đạt").WillReturnResult(sqlmock.NewResult(3, 1))

	// create add request
	addReq := &service.AddReq{
		Name: "Đạt",
	}
	// add customer
	_, err := s.AddCustomer(context.Background(), addReq)
	if err != nil {
		t.Errorf("query failed: %s", err.Error())
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// TestRetrieveCustomerCustomers tests logic of the RetrieveCustomers function
func TestRetrieveCustomer(t *testing.T) {

	// initialize connection
	initSet := initConn()
	s, mock := initSet()
	defer s.Db.Close()

	rows := sqlmock.NewRows([]string{"CustomerID", "Name"}).
		AddRow(1, "Son").
		AddRow(2, "Dat")

	mock.ExpectQuery("SELECT (.+) FROM Customer WHERE").WithArgs(2).WillReturnRows(rows)

	customer, err := s.RetrieveCustomer(context.Background(), &service.GetReq{
		Id: 2,
	})
	if err != nil {
		t.Errorf("query failed: %s", err)
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	fmt.Printf("Retrieved customer: id %d, name %v\n", customer.Id, customer.Name)
}

// TestDeleteCustomers tests logic of the DeleteCustomers function
func TestDeleteCustomer(t *testing.T) {
	// initialize connection
	initSet := initConn()
	s, mock := initSet()
	defer s.Db.Close()

	sqlmock.NewRows([]string{"CustomerID", "Name"}).
		AddRow(1, "Son").
		AddRow(2, "Dat")

	mock.ExpectExec("DELETE FROM Customer WHERE CustomerID").WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))

	_, err := s.DeleteCustomer(context.Background(), &service.DeleteReq{
		Id: 1,
	})
	if err != nil {
		t.Errorf("query failed: %s", err)
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
