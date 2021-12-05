package news

import (
	"os"
	"testing"
	"time"
)


func createConfig(t testing.TB, backupfile string, publishfile string) *PublisherConfig {
	t.Helper()
	return &PublisherConfig{
		Jsonformat: 	false,
		Backupfile: 	backupfile,
		Publishfile: 	publishfile,
	}
}

func TestPublisherStart(t *testing.T) {
	testCases := []struct {
		desc	string
		config  *PublisherConfig
	}{
		{desc: "start publisher", config: createConfig(t, "./test.txt", "./test.out")},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {

			p, err := Start("./test.txt")

			if err != nil || p.stopped != false {
				t.Fatalf("error starting publisher")
			}

			// if a Publisher is created during the test run, clean it up
			if p != nil && ! p.Stopped() {
				p.Stop()
			}
		})
	}
}

func TestPublisherStartFromStatefile(t *testing.T) {
	tc := struct {
		desc	string
		config 	*PublisherConfig
	}{
		desc: "test publisher start from statefile", 
		config: createConfig(t, "./test.txt", "./test.out"),
	}
	t.Run(tc.desc, func(t *testing.T) {

		p := NewPublisher(*tc.config)
		sub1 := NewSubscriber()
		sub2 := NewSubscriber()
		src := NewRandomSource()
		
		p.DistributeSource(src)
		
		sub1.Subscribe(p, "World")
		sub1.Subscribe(p, "Economics")
		sub2.Subscribe(p, "cooking")
		src.Publish()

		// TODO: this sleep is necessary ATM to avoid a race condition. If the publisher
		// gets cancelled, the source cannot be cancelled yet, this is still to be implemented.
		// Then, the situation can occur that the source publishes on a closed channel, causing
		// a panic. This sleeps is a workaround for that now. 
		time.Sleep(1 * time.Second)
		p.Stop()

		np, err := Start(tc.config.Backupfile)
		if err != nil {
			t.Fatalf("error starting publisher from configfile")
		}

		// check properties of new publisher
		act := np.config.Backupfile
		exp := tc.config.Backupfile
		if exp != act {
			t.Fatalf("error for config Backupfile, expected %s, got %s", exp, act)
		}

		act = np.config.Publishfile
		exp = tc.config.Publishfile
		if exp != act {
			t.Fatalf("error for config Publishfile, expected %s, got %s", exp, act)
		}

		bact := np.config.Jsonformat
		bexp := tc.config.Jsonformat
		if bexp != bact {
			t.Fatalf("error for config Jsonformat, expected %t, got %t", bexp, bact)
		}
		np.Stop()
	})
}

// This test is working for an empty publisher.
// TODO: issues with categories and articles, they are not properly marshalled into JSON.
func TestPublisherSave(t *testing.T) {
	testCases := []struct {
		desc	string
		config  *PublisherConfig
	}{
		{desc: "save publisher, correct path", config: createConfig(t, "./test.txt", "./test.out")},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			p := NewPublisher(*tc.config)

			p.Save()

			act, err := os.ReadFile(tc.config.Backupfile)
			if err != nil {
				t.Fatal(err)
			}

			exp := "{\"Config\":{\"Backupfile\":\"./test.txt\",\"Publishfile\":\"./test.out\",\"Jsonformat\":false},\"Categories\":[],\"Stopped\":false,\"Articles\":null}"
			if exp != string(act) {
				t.Fatalf("expected %s, got %s", exp, act)
			}

			// if a Publisher is created during the test run, clean it up
			if p != nil && ! p.Stopped() {
				p.Stop()
			}
		})
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

			p.articles = []Article{
				{id: 0, category: "world", title: "Ethiopia's Tigray conflict: Lalibela retaken", content: "Ethiopian troops have recaptured the historic town of Lalibela from Tigrayan rebels, the government has said."},
				{id: 1, category: "sports", title: "Max Verstappen & Lewis Hamilton set for thrilling Formula 1 finale", content: "The most intense Formula 1 championship fight for years will be decided over the next two weekends in the Middle East."},
				{id: 2, category: "local", title: "Clear Flour Bread serves the best cookies in Mass., according to Yelp", content: "It’s officially holiday cookie season, that time of year when Christmas tree-shaped sugar cookies lay snugly in tins next to snickerdoodles and gingersnap cookies."},
			}

			if p.Articles(1)[0].title != p.articles[1].title {
				t.Fatalf("error retrieving single article")
			}

			if p.Articles(1, 2)[0].title != p.articles[1].title {
				t.Fatalf("error retrieving multiple articles")
			}

			// if a Publisher is created during the test run, clean it up
			if p != nil && ! p.Stopped() {
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

			p.articles = []Article{
				{category: "world", title: "Ethiopia's Tigray conflict: Lalibela retaken", content: "Ethiopian troops have recaptured the historic town of Lalibela from Tigrayan rebels, the government has said."},
				{category: "sports", title: "Max Verstappen & Lewis Hamilton set for thrilling Formula 1 finale", content: "The most intense Formula 1 championship fight for years will be decided over the next two weekends in the Middle East."},
			}

			if len(p.articles) != 2 {
				t.Fatalf("error filling articles into archive of publisher")
			}

			p.Clear()

			if len(p.articles) != 0 {
				t.Fatalf("error clearing publisher archive")
			}

			// if a Publisher is created during the test run, clean it up
			if p != nil && ! p.Stopped() {
				p.Stop()
			}
		})
	}
}

