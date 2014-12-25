package errors

import "fmt"

/* Errors */

type NotFoundError struct {
	//Err     error
	Message string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%v", e.Message)
}
