package main

import (
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

type myHandler struct{}

func (mh myHandler) ServeHTTP(write http.ResponseWriter, request *http.Request) {

}
