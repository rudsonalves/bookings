package main

import "testing"

func TestRub(t *testing.T) {
	err := run()
	if err != nil {
		t.Errorf("fail run() with error: %s", err)
	}
}
