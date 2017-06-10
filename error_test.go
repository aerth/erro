package erro

import (
	"errors"
	"fmt"
	"io"
	"testing"
)

func TestEquality(t *testing.T) {
	if EOF == io.EOF {
		println("would be nice")
		println("this fail marks the end of this package")
		t.FailNow()

	}
	var myeof = io.EOF
	if myeof == io.EOF {
		println("myeof := io.EOF")
	}

	myeof2 := errors.New("EOF")
	if myeof2 == io.EOF {
		println("myeof2 := io.EOF")
	}

	if EOF.IsEqual(io.EOF) {
		println("EOF == io.EOF")
	}
}

func TestNewErr(t *testing.T) {
	f := func(err error) {
		if err == nil {
			fmt.Println("no error")
			return
		}
		fmt.Printf("%T %s\n", err, err)
	}
	e := Error("im not a string")
	f(&e) // take address for to implement error type

	// New creates a new Error and returns the pointer address
	err := New("i am a new error")
	if err != nil {
		fmt.Printf("%T %s\n", err, err)
	} else {
		panic("unreachable")
	}
}

func TestNilError(t *testing.T) {
	nilerr := func() error {
		return nil
	}
	nonnilerr := func() error {
		return New("hi")
	}

	err := nilerr()

	err = nonnilerr()
	fmt.Printf("%T %s\n", err, err)

}
