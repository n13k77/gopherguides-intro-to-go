package week04

import (
	"fmt"
	"io"
)

type Entertainer interface {
	Name() string
	Perform(v Venue) error
}

type Setuper interface {
	Setup(v Venue) error
}

type Teardowner interface {
	Teardown(v Venue) error
}

type messyEntertainer struct {
	FullName	string
}

func (m messyEntertainer) Name() string {
	return m.FullName
}

func (m messyEntertainer) Perform(v Venue) error {

	if v.Audience < 100 {
		return fmt.Errorf("entertainer %s refusing to perform for %d people, not worth the mess", m.FullName, v.Audience)
	}

	return nil
}

func (m messyEntertainer) Setup(v Venue) error {

	if v.Audience < 100 {
		return fmt.Errorf("entertainer %s refusing to set up for %d people, not worth the mess", m.FullName, v.Audience)
	}

	return nil
}

type cleanEntertainer struct {
	FullName	string
}

func (c cleanEntertainer) Name() string {
	return c.FullName	
}

func (c cleanEntertainer) Perform(v Venue) error {

	if v.Audience > 100 {
		return fmt.Errorf("entertainer %s refusing to perform for %d people, it will be too big of a mess", c.FullName, v.Audience)
	}

	return nil
}

func (c cleanEntertainer) Teardown(v Venue) error {

	if v.Audience > 100 {
		return fmt.Errorf("entertainer %s refusing tear down in front of for %d people, it's too big of a mess", c.FullName, v.Audience)
	}

	return nil
}

type Venue struct {
	Audience int
	Log      io.Writer
}

func (v Venue) show (ent Entertainer) error {

	name := ent.Name()

	if ent, ok := ent.(Setuper); ok {
		err := ent.Setup(v)
		if err != nil {
			return err
		}
		fmt.Fprintf(v.Log, "%s has completed setup.\n", name)
	}

	err := ent.Perform(v)
	if err != nil {
		return err
	}
	fmt.Fprintf(v.Log, "%s has performed for %d people.\n", name, v.Audience)

	if ent, ok := ent.(Teardowner); ok {
		err := ent.Teardown(v)
		if err != nil {
			return err
		}
		fmt.Fprintf(v.Log, "%s has completed teardown.\n", name)
	}

	return nil
}

func (v *Venue) Entertain (n int, e ...Entertainer) error {

	if len(e) == 0 {
		return fmt.Errorf("no entertainers to perform")
	}

	v.Audience = n
	for _, ent := range e {
		err := v.show(ent)
		if err != nil {
			return(err)
		}
	}

	return nil

} 
