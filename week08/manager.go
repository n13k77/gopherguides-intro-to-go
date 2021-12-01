package week08

import (
	"context"
	"sync"
)

// Manager is responsible for receiving product orders
// and assigning them to employees. Manager is also responsible
// for receiving completed products, and listening for errors,
// from employees. Manager takes products that have been built
// by employees and returns them to the customer as a CompletedProduct.
type Manager struct {
	Warehouse *Warehouse
	cancel    context.CancelFunc
	completed chan CompletedProduct
	errs      chan error
	jobs      chan *Product
	stopped   bool
	sync.RWMutex
}

// Start will create new employees for the given count,
// and start listening for jobs and errors.
// Managers should be stopped using the Stop method
// when they are no longer needed.
func (m *Manager) Start(ctx context.Context, count int) (context.Context, error) {

	if count <= 0 {
		return nil, ErrInvalidEmployeeCount(count)
	}

	// create a new cancellation context
	ctx, cancel := context.WithCancel(ctx)

	// hold onto the cancel function so it can be called
	// by m.Stop()
	m.cancel = cancel
	if m.Warehouse == nil {
		m.Warehouse = &Warehouse{}
	}

	// launch a goroutine to listen context cancellation
	go func(ctx context.Context) {

		// listen for context cancellation
		// this could come from the external context
		// passed to m.Start()
		<-ctx.Done()

		// call the cancel function
		cancel()

		// call Stop()
		m.Stop()
	}(ctx)

	// start the warehouse
	// this returns a context that can be listened to
	// for cancellation notification from the warehouse
	m.Lock()
	ctx = m.Warehouse.Start(ctx)
	m.Unlock()

	for i := 0; i < count; i++ {

		e := Employee(i + 1)

		// start the employee working
		// with the given context and manager
		go e.work(ctx, m)
	}

	// return the context for clients to listen to
	// for cancellation.
	return ctx, nil
}

// Assign will assign the given products to employees
// as employeess become available. An invalid product
// will return an error.
func (m *Manager) Assign(products ...*Product) error {
	if m.stopped {
		return ErrManagerStopped{}
	}

	// loop through each product and assign it to an employee
	for _, p := range products {
		// validate product
		if err := p.IsValid(); err != nil {
			return err
		}

		// assign product to employee
		// this will block until an employee becomes available
		m.Jobs() <- p
	}

	return nil
}

// Complete will wrap the employee and the product into
// a CompletedProduct. The will be passed down the Completed()
// channel as soon as a listener is available to receive it.
// Complete will error if the employee is invalid or
// if the product is not built.
func (m *Manager) Complete(e Employee, p *Product) error {
	// validate employee
	if err := e.IsValid(); err != nil {
		return err
	}

	// validate product is built
	if err := p.IsBuilt(); err != nil {
		return err
	}

	cp := CompletedProduct{
		Employee: e,
		Product:  *p, // deference pointer to value type product
	}

	// fmt.Printf("TODO >> manager.go:102 cp %[1]T %[1]v\n", cp)
	// Send completed product to Completed() channel
	// for a listener to receive it.
	// This will block until a listener is available.
	m.completedCh() <- cp

	return nil
}

// completedCh returns the channel for CompletedProducts
func (m *Manager) completedCh() chan CompletedProduct {
	m.Lock()
	defer m.Unlock()
	if m.completed == nil {
		m.completed = make(chan CompletedProduct)
	}
	return m.completed
}

// Completed will return a channel that can be listened to
// for CompletedProducts.
// This is a read-only channel.
func (m *Manager) Completed() <-chan CompletedProduct {
	return m.completedCh()
}

// Jobs will return a channel that can be listened to
// for new products to be built.
func (m *Manager) Jobs() chan *Product {
	m.Lock()
	defer m.Unlock()
	if m.jobs == nil {
		m.jobs = make(chan *Product)
	}
	return m.jobs
}

// Errors will return a channel that can be listened to
// and can be used to receive errors from employees.
func (m *Manager) Errors() chan error {
	m.Lock()
	defer m.Unlock()
	if m.errs == nil {
		m.errs = make(chan error)
	}
	return m.errs
}

// Stop will stop the manager and clean up all resources.
func (m *Manager) Stop() {
	m.Lock()
	defer m.Unlock()

	m.cancel()
	if m.stopped {
		return
	}

	m.stopped = true

	// close all channels
	if m.jobs != nil {
		close(m.jobs)
	}

	if m.errs != nil {
		close(m.errs)
	}

	if m.completed != nil {
		close(m.completed)
	}
}
