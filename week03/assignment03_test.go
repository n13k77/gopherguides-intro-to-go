package movies

import (
	"fmt"
	"math/rand"
	"testing"
)

// 	testing publicly accessible fields of the struct. This might look like overkill, but it
//   	turned out to be very valid in making sure the testing mechanism and referencing to the
//	movies package actually  worked.
func TestMovie(t *testing.T) {
	t.Parallel()
	_ = Movie{
		Length: 120,
		Name:   "testMovie",
	}
}

// test the Rate function in Error conditions
func TestRateError(t *testing.T) {
	t.Parallel()
	type testCase struct {
		plays int
	}
	testCases := []testCase{
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
	type testCase struct {
		plays   int
		viewers int
	}
	testCases := []testCase{
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
		if m.Plays() != tc.plays {
			t.Errorf("Expected %d plays, got %d", tc.plays, m.Plays())
		}
	}
}

// test Viewers function, makes use of the Play() function
func TestViewers(t *testing.T) {
	t.Parallel()
	type testCase struct {
		plays   int
		viewers int
	}
	testCases := []testCase{
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
		if m.Viewers() != (tc.plays * tc.viewers) {
			t.Errorf("Expected %d viewers, got %d", tc.plays*tc.viewers, m.Viewers())
		}
	}
}

// 	tests both the Rate and Rating function. The test package uses only methods to access the
//	private fields; movie rate can only be set with Rate function and read with Rating function
//	so these functions can only be tested in tandem.
func TestRating(t *testing.T) {
	t.Parallel()
	type testCase struct {
		plays  int
		rating float32
	}
	testCases := []testCase{
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
		if float32(m.Rating()) != tc.rating {
			t.Errorf("Expected %.2f, got %.2f", tc.rating, m.Rating())
		}

	}
}

// tests the String function
func TestString(t *testing.T) {
	t.Parallel()
	type testCase struct {
		title  string
		length int
		rate   float32
	}
	testCases := []testCase{
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
		expected := fmt.Sprintf("%s (%dm) %.2f%%", m.Name, m.Length, m.totalRating)
		if m.String() != expected {
			t.Errorf("Expected:\n %s\n, got:\n%s\n", expected, m.String())
		}
	}
}

// tests the Play function for Theatre struct
func TestTheatrePlay(t *testing.T) {
	t.Parallel()
	type testCase struct {
		movies  int
		viewers int
	}
	testCases := []testCase{
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
		theatre.Play(tc.viewers, m...)

		for i := 0; i < len(m); i++ {
			if m[i].viewers != tc.viewers {
				t.Errorf("Expected: %d viewers, got: %d", tc.viewers, m[i].viewers)
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

	testTheatre := Theatre{}
	var testMovies []*Movie

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
			testMovies = append(testMovies, &m)
		}
		err := testTheatre.Critique(testMovies, generateCritique)

		if err != nil {
			// I do not know how to wrap the received error here
			fmt.Println(err)
			t.Errorf("error while rating movie")
		}

		for i := 0; i < tc.numberOfMovies; i++ {
			if testMovies[i].Rating() <= 0 {
				t.Errorf("a movie did not get reviews")
			}
		}
	}
}
