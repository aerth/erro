package erro

// EOF can be compared to io.EOF by using the IsEqual method
const EOF = Error("EOF")

// Error is a simple error type that can easily be type converted from a string
type Error string

// New returns a new pointer to an Error
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
