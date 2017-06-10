package main

import (
	"errors"
	"io"
	"testing"
)

func TestEquality(t *testing.T) {
	if EOF == io.EOF {
		println("would be nice")
	}
	myeof := io.EOF
	if myeof == io.EOF {
		println("myeof := io.EOF")
	}
	myeof2 := errors.New("EOF")
	if myeof2 == io.EOF {
		println("myeof2 := io.EOF")

	}

	if EOF.IsEq(io.EOF) {
		println("EOF == io.EOF")

	}
}
