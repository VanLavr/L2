package middlewares

import (
	"time"

	"github.com/VanLavr/L2/develop/11/internal/pkg/errors"
)

// EventValidator (layout = "2006-01-02")
type EventValidator struct {
	layout string
}

// NewValidator is a constructor for validator
func NewValidator(layout string) *EventValidator {
	return &EventValidator{layout: layout}
}

// ValidateDate stands for validating date
func (v *EventValidator) ValidateDate(date string) (time.Time, error) {
	return time.Parse(v.layout, date)
}

// ValidateEventName stands for validating name of the event
func (v *EventValidator) ValidateEventName(name string) error {
	if name == "" {
		return errors.ErrProvidedEventNameIsInvalid
	}

	return nil
}
