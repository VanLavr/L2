package errors

import "errors"

var (
	ErrProvidedDateIsInvalid      = errors.New("provided date is invalid")
	ErrProvidedEventNameIsInvalid = errors.New("provided event name is invalid")
	ErrInternal                   = errors.New("internal server error")
	ErrNotFound                   = errors.New("provided was event not found")
)
