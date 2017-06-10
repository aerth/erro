package main

import (
	"io"
)

const EOF = Error("EOF")

type Error string

func New(s string) Error {
	return Error(s)
}

// stringer interface
func (e Error) String() string {
	return string(e)
}

// error interface
func (e Error) Error() string {
	return string(e)
}

// IsEq compares an error with e
func (e Error) IsEq(err error) bool {
	return err.Error() == string(e)
}

func main() {
	if EOF == io.EOF {
		println("would be nice")
	}
	myeof := io.EOF
	if myeof == io.EOF {
		println("myeof := io.EOF")
	}
	if EOF.IsEq(io.EOF) {
		println("EOF == io.EOF")

	}
	println("Hi")
}
