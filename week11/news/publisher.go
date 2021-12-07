package news

import (
	"log"
	"strings"
	"sync"
)

type Publisher struct {
	config 		PublisherConfig
	mutex  		sync.RWMutex
	cats		[][]string
	subs 		[]chan Article
	stopped		bool
	archive		[]Article
}

type PublisherConfig struct {
	Backupfile 	string
	Publishfile string
	Jsonformat  bool
}

// func NewPublisher() creates a new instance of a Publisher
func NewPublisher(config PublisherConfig) (*Publisher) {
	log.Println("publisher new publisher")
	p := &Publisher{}

	p.subs = []chan Article{} 	// slice of channels, the slice index serves as the subscriber ID
	p.cats = [][]string{}		// slice of slice of categories, the slice index serves as the subscriber ID
	p.stopped = false

	// if not specified, set backupfile to a default location
	if config.Backupfile == "" {
		config.Backupfile = "./backup.tmp"
	}

	p.config = config

	return p
}

// func AddSource() adds a news source to the publisher, distributes its articles
func (p *Publisher) Dispatch(s Source) {
	log.Println("publisher distribute source")

	go func() {
		// start listening to articles that are published by the source
		log.Println("publisher distribute source start listening")
		for a := range s.GetSourceChannel() {

			// add the received article to the archive 
			p.mutex.Lock()
			a.id = len(p.archive) + 1
			p.archive = append(p.archive, a)
			p.mutex.Unlock()

			lc := strings.ToLower(a.Category())

			for i, topics := range p.cats {
				for _, topic := range topics {
					if topic == lc {
						p.subs[i] <- a
					}
				}
			}
				
		}
	}()
}

func (p *Publisher)Subscribe(cats []string) (<-chan Article, error) {
	log.Println("publisher subscribe")
	p.mutex.Lock()
	defer p.mutex.Unlock()

	ch := make(chan Article)
	id := len(p.subs)
	
	for _, cat := range cats {
		p.cats[id] = append(p.cats[id],strings.ToLower(cat))
	} 
	p.subs[id] = ch
	
	return ch, nil
}

func (p *Publisher) Stopped() bool {
	log.Println("publisher stopped")
	return p.stopped
}

func (p *Publisher) Articles (ids ...int) []Article {
	log.Printf("publisher articles")

	p.mutex.RLock()
	defer p.mutex.RUnlock()
	a := []Article{}

	for _, id := range ids {
		log.Printf("publisher articles id %d", id)
		a = append(a, p.archive[id])	
	}
	return a
}

// func Clear clears the publisher archive
func (p *Publisher) Clear () {
	log.Printf("publisher clear")
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.archive = nil
}

