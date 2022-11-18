package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH myHandler

	h := NoSurf(myH)

	switch h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Error("type is not http.Handler")
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler

	h := SessionLoad(myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("type is not http.Handler, type is %T", v)
	}
}
