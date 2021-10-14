package movies

import (
	_ "fmt"
)

/* 	Define a Movie struct and export two public fields; Length of type int and
Name of type string. You must NOT export any more public fields on Movie.
You are allowed, however, to use non-exported, private fields as needed.
*/

type Movie struct {
	Lenght int
	Name   string
}

/* 	Define a new type named CritiqueFn that is a function type that takes a
pointer of type Movie and returns a float32 and an error. Note: This is not
a func definition, but rather a type definition.
*/

type CritiqueFn func(m *Movie) (float32, error)

/*	Define a Theatre struct. You must not export any public fields on Theatre.
	You are allowed, however, to use non-exported, private fields as needed.
*/

type Theatre struct {
	a int // TODO: remove, dummy declaration
}

/* 	Define a method on the pointer receiver for Movie named Rate. Rate should
take a float32 rating and return an error. Calling Rate should track this
rating. If the number of plays is 0 return the following error:
fmt.Errorf("can't review a movie without watching it first")
*/

func (m *Movie) Rate(rating float32) error {
	return errors.New("TODO replace message")
}

/* 	Define a method on the pointer receiver for Movie named Play. Play should
take the number (int) of viewers watching the movie. Calling Play should
increase both the number of viewers, as well as the number of plays, for
the movie.
*/

func (m *Movie) Play(viewers int) {
	_
}

/* 	Define a method on the value receiver for Movie named Viewers. Viewers takes
no arguments and returns the number (int) of people who have viewed the movie.
*/

func (m Movie) Viewers () int {
	return 0
}

/* 	Define a method on the value receiver for Movie named Plays. Plays takes no
arguments and returns the number (int) of times the movie has been played.
*/

func (m Movie) Plays () int {
	return 0
}

/*	Define a method on the value receiver for Movie named Rating. Rating takes no
	arguments and returns the rating (float64) of the movie. This can be calculated
	by the total ratings for the movie divided by the number of times the movie has
	been played.
*/

func (m Movie) Rating () int {
	return 0
}

/*	Define a method on the value receiver for Movie named String. String should
	return a string that that includes the name, length, and rating of the film.
	Ex. Wizard of Oz (102m) 99.0%
*/

func (m Movie) String () string {
	return "a"
}

/*	Define a method on the pointer receiver for Theatre named Play. Play should
	take the number (int) of viewers, a variadic list of pointer of type Movie,
	and return an error. Calling Play will call each movie’s Play method with the
	number of viewers. If no movies are passed in return the following error:
	fmt.Errorf("no movies to play").
*/

func (t *Theatre) Play (viewers int, movies ...*Movie) (string, error) {
	return "a", nil
}

/*	Define a method on the value receiver for Theatre named Critique. Critique
	should take a slice of pointers to type Movie, a CritiqueFn, and return an
	error. Calling Critique will iterate over the movies and for each movie.
	First, each movie’s Play method should be called with a value of 1. Next,
	the CritiqueFn should be called with the movie, the return values should
	be error checked. If there is no error the movie’s Rate method should be
	called with the float32 value that was returned from the CritiqueFn. Again,
	this call should be error checked.
*/

func (t Theatre) Critique (movies []*Movie, fn CritiqueFn) error {
	return nil
}
