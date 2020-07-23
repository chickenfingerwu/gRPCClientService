package main

import "database/sql"

type SQLDB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type MockDB struct {
	callParam []interface{}
}

// Implement SQLDB interface
func (mdb *MockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	mdb.callParam = []interface{}{query}
	mdb.callParam = append(mdb.callParam, args...)

	return nil, nil
}
