package demo

import (
	"errors"
	"fmt"
)

type ErrTableNotFound struct {
	table string
}

func (e ErrTableNotFound) Error() string {
	return fmt.Sprintf("table not found %s", e.table)
}

func (e ErrTableNotFound) TableNotFound() string {
	return e.table
}

func (e ErrTableNotFound) Is(err error) bool {
	_, ok := err.(ErrTableNotFound)
	return ok
}

func IsErrTableNotFound(err error) bool {
	return errors.Is(err, ErrTableNotFound{})
}

// --- ErrNoRows ---

type ErrNoRows interface {
	error
	RowNotFound() (string, Clauses)
}

var _ ErrNoRows = &errNoRows{} // asserting interface implementation

type errNoRows struct {
	clauses Clauses
	table   string
}

func (e *errNoRows) Error() string {
	return fmt.Sprintf("[%s] no rows found\nquery: %s", e.table, e.Clauses())
}

func (e *errNoRows) Clauses() Clauses {
	if e.clauses == nil {
		e.clauses = Clauses{}
	}

	return e.clauses
}

func (e *errNoRows) RowNotFound() (string, Clauses) {
	return e.table, e.Clauses()
}

func (e *errNoRows) Is(err error) bool {
	_, ok := err.(*errNoRows)
	return ok
}

func IsErrNoRows(err error) bool {
	return errors.Is(err, &errNoRows{})
}

func AsErrNoRows(err error) (ErrNoRows, bool) {
	e := &errNoRows{}
	if errors.As(err, &e) {
		return e, true
	}

	return nil, false
}
