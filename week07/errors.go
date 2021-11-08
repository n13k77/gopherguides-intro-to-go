package week07

import "fmt"

// ErrInvalidQuantity is returned when the product quantity is invalid.
type ErrInvalidQuantity int

func (e ErrInvalidQuantity) Error() string {
	return fmt.Sprintf("quantity must be greater than 0, got %d", e)
}

// ---

// ErrProductNotBuilt is returned when the product is not built.
type ErrProductNotBuilt string

func (e ErrProductNotBuilt) Error() string {
	return string(e)
}

// ---

// ErrInvalidEmployee is returned when the employee number is invalid.
type ErrInvalidEmployee int

func (e ErrInvalidEmployee) Error() string {
	return fmt.Sprintf("invalid employee number: %d", e)
}

// ---

// ErrInvalidEmployeeCount is returned when the employee count is invalid.
type ErrInvalidEmployeeCount int

func (e ErrInvalidEmployeeCount) Error() string {
	return fmt.Sprintf("invalid employee count: %d", e)
}

// ---

type ErrManagerStopped struct{}

func (ErrManagerStopped) Error() string {
	return "manager is stopped"
}
