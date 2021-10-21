package movies

import (
	"fmt"
	"math/rand"
	"testing"
)

// test the Rate function in Error conditions
func TestRateError(t *testing.T) {
	t.Parallel()

	testCases := []struct{
		plays int
	}{
		{plays: 0},
		{plays: 10},
		{plays: 100},
	}

	for _, tc := range testCases {

		m := Movie{
			Name:   "TestTitle",
			Length: 10,
		}

		for i := 0; i < tc.plays; i++ {
			m.Play(1)
			err := m.Rate(1.0)
			if tc.plays == 0 && err == nil {
				t.Errorf("expected error for rating with 0 plays, got no error")
			}
			if tc.plays != 0 && err != nil {
				t.Errorf("was not expecting error for rating with >0 plays, got error")
			}
		}
	}
}

// 	tests both the Play and Plays function for Movie struct. The test package uses only
//	methods to access the private fields; movie play can only be set with Play function
//	and read with Plays function so these functions can only be tested in tandem. Makes
//	use of Viewers() function
func TestPlay(t *testing.T) {

	t.Parallel()

	testCases := []struct{
		plays   int
		viewers int
	}{
		{plays: 0, viewers: 10},
		{plays: 20, viewers: 30},
		{plays: 5000, viewers: 20},
	}

	for _, tc := range testCases {

		m := Movie{
			Name:   "TestTitle",
			Length: 0,
		}

		for i := 0; i < tc.plays; i++ {
			m.Play(tc.viewers)
		}

		exp := m.Plays()
		act := tc.plays

		if exp != act {
			t.Errorf("expected %d plays, got %d", exp, act)
		}
	}
}

// test Viewers function, makes use of the Play() function
func TestViewers(t *testing.T) {

	t.Parallel()

	testCases := []struct{
		plays   int
		viewers int
	}{
		{plays: 0, viewers: 10},
		{plays: 20, viewers: 30},
		{plays: 5000, viewers: 20},
	}

	for _, tc := range testCases {

		m := Movie{
			Name:   "TestTitle",
			Length: 0,
		}

		for i := 0; i < tc.plays; i++ {
			m.Play(tc.viewers)
		}

		act := m.Viewers()
		exp := tc.plays * tc.viewers

		if act != exp {
			t.Errorf("expected %d viewers, got %d", exp, act)
		}
	}
}

// 	tests both the Rate and Rating function. The test package uses only methods to access the
//	private fields; movie rate can only be set with Rate function and read with Rating function
//	so these functions can only be tested in tandem.
func TestRating(t *testing.T) {

	t.Parallel()

	testCases := []struct{
		plays  int
		rating float32
	}{
		{plays: 1, rating: 1.0},
		{plays: 10, rating: 5.0},
		{plays: 10, rating: 0.0},
		{plays: 5000, rating: 10.0},
	}

	for _, tc := range testCases {

		m := Movie{
			Name:   "testTitle",
			Length: 0,
		}

		for i := 0; i < tc.plays; i++ {
			m.Play(1)
			m.Rate(tc.rating)
		}
		// Rate and Rating use float32 and float64, respectively. Conversion needed to compare

		act := float32(m.Rating())
		exp := tc.rating

		if  act != exp {
			t.Errorf("expected %.2f, got %.2f", exp, act)
		}
	}
}

// tests the String function
func TestString(t *testing.T) {

	t.Parallel()

	testCases := []struct{
		title  string
		length int
		rate   float32
	}{
		{"TestTitle 1", 100, 0.0},
		{"A 'slightly' more advanced _ title!", 20, 8.0},
		{"你好吗", 75, 6.5},
	}

	for _, tc := range testCases {

		m := Movie{
			Name:   tc.title,
			Length: tc.length,
		}

		m.Rate(tc.rate)
		act := m.String()
		exp := fmt.Sprintf("%s (%dm) %.2f%%", m.Name, m.Length, m.totalRating)

		if act != exp {
			t.Errorf("expected:\n %s\n, got:\n%s\n", exp, act)
		}
	}
}

// tests the Play function for Theatre struct
func TestTheatrePlay(t *testing.T) {

	t.Parallel()

	testCases := []struct{
		movies  int
		viewers int
	}{
		{2, 300},
		{4, 75},
		{10, 100},
	}

	for _, tc := range testCases {
		theatre := Theatre{}
		m := []*Movie{}
		for i := 0; i < tc.movies; i++ {
			testCase := Movie{
				Name:   "TestTitle",
				Length: 10,
			}
			m = append(m, &testCase)
		}

		err := theatre.Play(tc.viewers, m...)
		if err != nil {
			t.Error(err)
		}

		for i := 0; i < len(m); i++ {

			act := m[i].viewers
			exp := tc.viewers

			if act != exp {
				t.Errorf("expected: %d viewers, got: %d", exp, act)
			}
		}
	}
}

// test Critique function
func TestCritique(t *testing.T) {

	t.Parallel()

	var generateCritique CritiqueFn = func(m *Movie) (float32, error) {
		if m.viewers < 10 {
			return 0.0, fmt.Errorf("not enough viewers to count the rating")
		}
		return rand.Float32() * 100.0, nil
	}

	theatre := Theatre{}
	var movies []*Movie

	type testCase struct {
		numberOfMovies int
		viewers        int
	}
	testCases := []testCase{
		{numberOfMovies: 2, viewers: 100},
		{numberOfMovies: 50, viewers: 50},
		{numberOfMovies: 100, viewers: 20},
	}
	for _, tc := range testCases {
		for i := 0; i < tc.numberOfMovies; i++ {
			m := Movie{
				Name:   "TestName",
				Length: 10,
			}
			m.Play(tc.viewers)
			movies = append(movies, &m)
		}
		err := theatre.Critique(movies, generateCritique)

		if err != nil {
			t.Error(err)
		}

		for i := 0; i < tc.numberOfMovies; i++ {
			if movies[i].Rating() <= 0 {
				t.Errorf("a movie did not get reviews")
			}
		}
	}
}
