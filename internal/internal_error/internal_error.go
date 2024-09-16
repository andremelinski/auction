package internalerror

type InternalError struct {
	Message string
	Err     string
}

func (ie *InternalError) Error() string {
	return ie.Message
}

func NewNotFoundError(message string) *InternalError {
	return &InternalError{
		Message: message,
		Err:     "not_found",
	}
}

func NewInternalserverError(message string) *InternalError {
	return &InternalError{
		Message: message,
		Err:     "internal_server_error",
	}
}

/*
The Error() function in Go is part of the error interface
By implementing the Error() method in your InternalError struct, you allow instances of InternalError to be treated as
errors throughout your Go application. This is useful because any type that implements the Error() method is recognized as an error type by Go.

Why implement the Error() function:
Standardization: It allows your custom error type (InternalError) to conform to Go’s error interface.
	This means your custom errors can be returned from functions that expect errors and passed around as part of Go's error-handling mechanisms.

Custom error messages: By implementing Error(), you can customize what message gets returned when Error() is called. In your case, Error() returns ie.Message,
	so whenever this error is printed or logged, the Message field is shown.

Compatibility: Since Go functions that deal with errors (like logging functions or error-checking utilities) expect a type that implements the error interface,
	your custom error type will integrate smoothly with existing error-handling code.

Example:
err := NewNotFoundError("resource not found")
fmt.Println(err)  // Output: resource not found
Here, the Error() method is automatically called to retrieve the string representation of the error (which in this case is the Message).

*/
