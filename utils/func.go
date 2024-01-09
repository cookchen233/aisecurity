package utils

import (
	"fmt"
	"github.com/pkg/errors"
)

func FilterZerosFromArray(input []int) []int {
	var result []int
	for _, value := range input {
		if value != 0 {
			result = append(result, value)
		}
	}
	return result
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// ErrorWrap adds a message to the error and ensures the stack trace is only added once.
func ErrorWrap(err error, message string) error {
	// Check if the error has a stack trace
	fmt.Println(err.Error())
	if _, ok := err.(stackTracer); ok {
		// The error already has a stack trace; add only the message.
		return errors.WithMessage(err, message)
	}

	// The error does not have a stack trace; add both message and stack trace.
	return errors.Wrap(err, message)
}

// ErrorWithStack adds a stack trace to the error if it doesn't have one already.
func ErrorWithStack(err error) error {
	// Check if the error has a stack trace
	fmt.Println(err.Error())
	if _, ok := err.(stackTracer); ok {
		// Error already has a stack trace, return it as is
		return err
	}

	// Error does not have a stack trace, add it
	return errors.WithStack(err)
}
