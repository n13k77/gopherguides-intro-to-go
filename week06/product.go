package week06

import (
	"fmt"
	"time"
)

// Product to be built by an employee
type Product struct {
	Quantity int

	// non-exported fields (PRIVATE)
	// !YOU MAY NOT ACCESS THESE FIELDS IN YOUR TESTS!
	builtBy Employee
}

// BuiltBy returns the employee that built the product.
// A return value of "0" means no employee has built the product yet.
func (p Product) BuiltBy() Employee {
	return p.builtBy
}

// Build builds the product by the given employee.
// Returns an error if the product has already been built.
// Returns an error if the employee ID <= 0.
// Returns an error if the quantity <= 0.
func (p *Product) Build(e Employee) error {
	// error check

	if err := p.IsValid(); err != nil {
		return err
	}

	if err := e.IsValid(); err != nil {
		return err
	}

	// build the product here
	time.Sleep(time.Millisecond * time.Duration(p.Quantity))

	// mark the product as built
	p.builtBy = e

	return nil
}

// IsValid returns an error if the product is invalid.
// A valid product has a quantity > 0.
func (p Product) IsValid() error {
	if p.Quantity <= 0 {
		return ErrInvalidQuantity(p.Quantity)
	}

	return nil
}

// IsBuilt returns an error if the product is not built,
// or if the product is invalid.
func (p Product) IsBuilt() error {
	if err := p.IsValid(); err != nil {
		return err
	}

	if p.builtBy == 0 {
		return ErrProductNotBuilt(fmt.Sprintf("product is not built: %v", p))
	}

	return nil
}
