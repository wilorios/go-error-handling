package error_type

import "fmt"

type Operation string
type Category string

type CustomTypeError struct {
	Op   Operation // operation
	Kind Category  //category of the error
	Err  error     //the wrapped error
	// application specific fields
}

// Error implements error
func (e *CustomTypeError) Error() string {
	return fmt.Sprintf("Kind %v: Operation %v Error %v", e.Op, e.Kind, e.Err)
}
