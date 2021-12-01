package week08

// CompletedProduct represents a completed product
type CompletedProduct struct {
	Product  Product  // Built Product
	Employee Employee // Employee who built the product
}

// IsValid returns true if the product has been built.
func (cp CompletedProduct) IsValid() error {
	if err := cp.Employee.IsValid(); err != nil {
		return err
	}

	if err := cp.Product.IsBuilt(); err != nil {
		return err
	}

	return nil
}
