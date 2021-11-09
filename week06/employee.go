package week06

// Employee is a worker.
type Employee int

// IsValid returns an error if the employee is not valid.
// A valid employee is greater than zero.
//	valid: Employee(1)
//	valid: Employee(2)
//	invalid: Employee(0) 
//	invalid: Employee(-1)
func (e Employee) IsValid() error {
	if e > 0 {
		return nil
	}

	return ErrInvalidEmployee(e)
}

// worker listens for work from the manager
// and tries to complete it.
func (e Employee) work(m *Manager) {

	// Use an infinite loop so we can listen for the next
	// message coming down a channel.
	// Without an infinite loop, the select statement
	// would process the first channel with a message
	// and then exit.
	for {

		// listen for messages on different channels
		select {
		case <-m.Done(): // listen for the manager to signal that it is done
			return
		case p, ok := <-m.Jobs(): // listen for a new job

			// check if the channel is closed or not
			if !ok {
				continue // skip the rest of the loop
			}

			// try to build the product
			err := p.Build(e)

			if err != nil {
				// if there is an error, send it to the manager
				m.Errors() <- err
				continue // skip the rest of the loop
			}

			// if there is no error, send the product back to the manager
			m.Complete(e, p)

		}
	}

}
