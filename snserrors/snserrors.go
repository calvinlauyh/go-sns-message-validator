package snserrors

// snsError is an private structure to record an SNS error
type SNSError struct {
	t string // Type
	s string // Message
}

// Create and return an error with given type and message
func New(errType string, errMsg string) *SNSError {
	return &SNSError{errType, errMsg}
}

// Return the type of the error
func (err *SNSError) Type() string {
	return err.t
}

// Determine if an error is the given type.
// Returns true if the error is exactly the given type, false otherwise
func (err *SNSError) Is(errType string) bool {
	return err.t == errType
}

// Return the message of the error. This method also control how fmt package
// formats the error value
func (err *SNSError) Error() string {
	return err.s
}
