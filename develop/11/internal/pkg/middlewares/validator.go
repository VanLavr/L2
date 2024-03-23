package middlewares

import (
	"time"

	"github.com/VanLavr/L2/develop/11/internal/pkg/errors"
)

// Validator (layout = "2006-01-02")
type Validator struct {
	layout string
}

// NewValidator is a constructor for validator
func NewValidator(layout string) *Validator {
	return &Validator{layout: layout}
}

// ValidateDate stands for validating date
func (v *Validator) ValidateDate(date string) (time.Time, error) {
	return time.Parse(v.layout, date)
}

// ValidateEventName stands for validating name of the event
func (v *Validator) ValidateEventName(name string) error {
	if name == "" {
		return errors.ErrProvidedEventNameIsInvalid
	}

	return nil
}
