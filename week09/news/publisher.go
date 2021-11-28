package news

import (
	"sync"
)

// type Publisher publishes news articles to which can be subscribed
type Publisher struct {
	mutex  	sync.RWMutex
	subs   	map[string]map[int]chan Article
	stopped	bool
	state	string // location of the state file 
}

type PublisherConfig struct {
	backupfile	string


}
// func NewPublisher() creates a new instance of a Publisher
func NewPublisher(state string) (*Publisher, error) {
	p := &Publisher{}
	p.subs = make(map[string]map[int]chan Article)
	p.stopped = false
	p.state = "" // TODO
	return p, nil // TODO, return an error in case the statefile cannot be written to
}

// func Subscriber() adds a subscriber to a publisher. The subscriber has to
// provide the topic to which it is subscribing and its unique identifier

// TODO: enable subscribing to multiple topics
// TODO: ensure that the subscriber is subscribing with a truly unique ID
// TODO: enable subscribing to multiple or all topics at once
func (p *Publisher) Subscribe(category string, id int) <-chan Article {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	ch := make(chan Article, 1)
	p.subs[category][id] = ch
	return ch
}

// func Publish publishes an article to the subscribers subscribed to that 
// particular topic
func (p *Publisher) Publish(a Article) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	if p.stopped {
		return
	}

	for subs := range p.subs[a.category] {
		p.subs[a.category][subs] <- a
	}
}

// func Close closes the publisher 
func (p *Publisher) Close() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if !p.stopped {
		p.stopped = true
		for _, subs := range p.subs {
			for _, ch := range subs {
				close(ch)
			}
		}
	}
}

// func Save saves the state of the publisher to the configured saving location 
func (p *Publisher) Save() {
	//TODO
}