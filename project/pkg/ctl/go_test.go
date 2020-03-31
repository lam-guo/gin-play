package ctl

import "testing"

var a string
var done bool

func step() {
	a = "test"
	done = true
}

func TestLine(t *testing.T) {
	go step()
	for !done {
	}
	t.Log(a)
}
