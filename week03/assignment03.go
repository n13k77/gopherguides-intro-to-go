package movies

import (
	"fmt"
)

// A Movie represents a movie that plays in a Theatre
type Movie struct {
	Length      int     // the length of a Movie, in minutes
	Name        string  // the name of the Movie
	plays       int     // the amount of times it was played
	viewers     int     // the viewers that have seen the Movie
	totalRating float64 // the total rating the Movie has recieved
}

// the signature of a function that can be implemented to generate a critique for a Movie
type CritiqueFn func(m *Movie) (float32, error)

// a Theatre represents a theatre were a movie can be played
type Theatre struct {
	// empty struct
}

// Rate can be used to give a rating to a movie. It returns an error when the movie has not been watched first
func (m *Movie) Rate(rating float32) error {
	if m.plays == 0 {
		return fmt.Errorf("can't review a movie without watching it first")
	}
	m.totalRating += float64(rating)
	return nil
}

// Play increases the amount of times that the Movie has been played by 1, and increases the amount of
// viewers of the movie with the particular amount that watched the Movie this time.
func (m *Movie) Play(viewers int) {
	m.viewers += viewers
	m.plays++
}

// Viewers returns the total amount of people that has watched the movie
func (m Movie) Viewers() int {
	return m.viewers
}

// Plays returns the total amount of time that the movie has been played
func (m Movie) Plays() int {
	return m.plays
}

// Rating returns the average rating that the movie has received per play
func (m Movie) Rating() float64 {
	if m.plays == 0 {
		return 0
	}
	return m.totalRating / float64(m.plays)
}

// String returns a string that that includes the name, length, and rating of the film
func (m Movie) String() string {
	return fmt.Sprintf("%s (%dm) %.2f%%", m.Name, m.Length, m.totalRating)
}

// Play takes the number of viewers, and a variadic list of pointer of type Movie, and return an error. Calling Play will call each movie’s Play method with the number of viewers. If no movies are passed in return the following error: fmt.Errorf("no movies to play").
func (t *Theatre) Play(viewers int, movies ...*Movie) error {
	if len(movies) == 0 {
		return fmt.Errorf("no movies to play")
	}

	for _, m := range movies {
		m.Play(viewers)
	}
	return nil
}

func (t Theatre) Critique(movies []*Movie, fn CritiqueFn) error {

	for _, m := range movies {
		m.Play(1)
		critique, err := fn(m)

		if err != nil {
			return fmt.Errorf("error while playing movie: %w", err)
		}

		err = m.Rate(critique)

		if err != nil {
			return fmt.Errorf("error while rating movie %w", err)
		}
	}

	return nil
}
