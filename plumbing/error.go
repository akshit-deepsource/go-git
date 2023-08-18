package plumbing

import (
	"fmt"

	"github.com/pkg/errors"
)

type PermanentError struct {
	Err error
}

func NewPermanentError(err error) error {
	if err == nil {
		return nil
	}

	return errors.WithStack(&PermanentError{Err: err})
}

func (e *PermanentError) Error() string {
	return fmt.Sprintf("permanent client error: %s", e.Err.Error())
}

type UnexpectedError struct {
	Err error
}

func NewUnexpectedError(err error) error {
	if err == nil {
		return nil
	}

	return errors.WithStack(&UnexpectedError{Err: err})
}

func (e *UnexpectedError) Error() string {
	return fmt.Sprintf("unexpected client error: %s", e.Err.Error())
}
