package main

import (
	"net/http"
	"os"
	"testing"
)

// Before starting running the tests, do something inside this func, then run the tests and exit.
func TestMain(m *testing.M) {

	os.Exit(m. Run())
}

type myHandler struct {}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
}