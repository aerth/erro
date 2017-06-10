package erro

const EOF = Error("EOF")

type Error string

func New(s string) *Error {
	var e = new(Error)
	*e = Error(s)
	return e
}

// stringer interface
func (e Error) String() string {
	return string(e)
}

// error interface
func (e Error) Error() string {
	return string(e)
}

// IsEqual compares an error with e
func (e Error) IsEqual(err error) bool {
	return err.Error() == string(e)
}
