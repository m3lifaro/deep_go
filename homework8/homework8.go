package homework8

import (
	"errors"
	"strconv"
	"strings"
)

type MultiError struct {
	errors []error
}

func (e *MultiError) Error() string {
	sb := new(strings.Builder)
	sb.WriteString(strconv.Itoa(len(e.errors)))
	sb.WriteString(" errors occured:\n")
	for _, err := range e.errors {
		sb.WriteString("\t* ")
		sb.WriteString(err.Error())
	}
	sb.WriteString("\n")
	return sb.String()
}

func Append(err error, errs ...error) *MultiError {
	var multiError *MultiError
	if errors.As(err, &multiError) {
		multiError, _ := err.(*MultiError)
		multiError.errors = append(multiError.errors, errs...)
		return multiError
	}

	return &MultiError{errors: errs}
}
