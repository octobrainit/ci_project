package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Soma result is invalid: Expected -> %d. Recieved -> %d", 30, total)
	}
}
