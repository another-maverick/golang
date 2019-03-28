package main

import "testing"

func TestName(t *testing.T) {
	name := foo()

	if name != "World!" {
		t.Error("foo had not returned World! Something wronf")
	}
}
