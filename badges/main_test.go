package main

import "testing"

func Test_add(t *testing.T) {
	c := add(1, 2)

	if c != 3 {
		t.Error("got error")
	}

	d := sub(2, 1)

	if d != 1 {
		t.Error("got error")
	}
}
