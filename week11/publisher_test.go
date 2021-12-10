package news

import (
	"testing"
)

func createConfig(t testing.TB, backupfile string, publishfile string) *PublisherConfig {
	t.Helper()
	return &PublisherConfig{
		Jsonformat: 	false,
		Backupfile: 	backupfile,
		Publishfile: 	publishfile,
	}
}


func TestPublisherArticles(t *testing.T) {
	testCases := []struct {
		desc	string
		config  *PublisherConfig
	}{
		{desc: "clear publisher", config: createConfig(t, "./test.txt", "./test.out")},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			p := NewPublisher(*tc.config)

			p.Archive = []Article{
				{Id: 0, Category: "world", Title: "Ethiopia's Tigray conflict: Lalibela retaken", Content: "Ethiopian troops have recaptured the historic town of Lalibela from Tigrayan rebels, the government has said."},
				{Id: 1, Category: "sports", Title: "Max Verstappen & Lewis Hamilton set for thrilling Formula 1 finale", Content: "The most intense Formula 1 championship fight for years will be decided over the next two weekends in the Middle East."},
				{Id: 2, Category: "local", Title: "Clear Flour Bread serves the best cookies in Mass., according to Yelp", Content: "It’s officially holiday cookie season, that time of year when Christmas tree-shaped sugar cookies lay snugly in tins next to snickerdoodles and gingersnap cookies."},
			}

			if p.Articles(1)[0].Title != p.Archive[1].Title {
				t.Fatalf("error retrieving single article")
			}

			if p.Articles(1, 2)[0].Title != p.Archive[1].Title {
				t.Fatalf("error retrieving multiple articles")
			}

			// if a Publisher is created during the test run, clean it up
			if p != nil && p.Stopped == false {
				p.Stop()
			}
		})
	}
}

func TestPublisherClear(t *testing.T) {
	testCases := []struct {
		desc	string
		config  *PublisherConfig
	}{
		{desc: "clear publisher", config: createConfig(t, "./test.txt", "./test.out")},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			p := NewPublisher(*tc.config)

			p.Archive = []Article{
				{Category: "world", Title: "Ethiopia's Tigray conflict: Lalibela retaken", Content: "Ethiopian troops have recaptured the historic town of Lalibela from Tigrayan rebels, the government has said."},
				{Category: "sports", Title: "Max Verstappen & Lewis Hamilton set for thrilling Formula 1 finale", Content: "The most intense Formula 1 championship fight for years will be decided over the next two weekends in the Middle East."},
			}

			if len(p.Archive) != 2 {
				t.Fatalf("error filling articles into archive of publisher")
			}

			p.Clear()

			if len(p.Archive) != 0 {
				t.Fatalf("error clearing publisher archive")
			}

			// if a Publisher is created during the test run, clean it up
			if p != nil && p.Stopped == false {
				p.Stop()
			}
		})
	}
}