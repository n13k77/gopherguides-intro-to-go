package news

import (
	"os"
	"sync"
)

// type Publisher publishes news articles to which can be subscribed
type Publisher struct {
	mutex  		sync.RWMutex
	subs   		map[string]map[int]chan Article
	stopped		bool
	backupfile 	*os.File
	publishfile	*os.File // this is part of optional functionality, check for nil before using
	jsonformat	bool
}

type PublisherConfig struct {
	backupfile	string
	publishfile	string
	jsonformat	bool
}

// func NewPublisher() creates a new instance of a Publisher
func NewPublisher(config PublisherConfig) (*Publisher, error) {
	var err error

	p := &Publisher{}
	p.subs = make(map[string]map[int]chan Article)
	p.stopped = false

	// if not specified, set backupfile to a default location
	if config.backupfile == "" {
		config.backupfile = "./tmp.txt"
	}

	// check writability of file locations
    p.backupfile, err = os.Create(config.backupfile)
    if err != nil {
        return nil, err
    }

	if config.publishfile != "" {
		p.jsonformat = config.jsonformat
		p.publishfile, err = os.Create(config.publishfile)
	}

	if err != nil {
        return nil, err
    }

    defer func() {
        if err := p.backupfile.Close(); err != nil {
            panic(err)
        }
    }()

	return p, nil
}

// func Subscribe() adds a subscriber to a publisher. The subscriber has to
// provide the topic to which it is subscribing and its unique identifier
func (p *Publisher) Subscribe(id int, categories ...string) (<-chan Article, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	ch := make(chan Article, 1)

	for _, category := range categories {
		p.subs[category][id] = ch
	}
	return ch, nil
}

// func Unsubscriber() removes a subscriber from a publisher. The subscriber has to
// provide the topic from which it is unsubscribing and its unique identifier
func (p *Publisher) Unsubscribe(id int, categories ...string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	cats := categories

	// if no categories are provided, remove all subscriptions for this subscriber
	if len(cats) == 0 {
		i := 0
		for k := range(p.subs) {
			cats[i] = k
			i++
		}
	}

	for _, cat := range cats {
		delete(p.subs[cat], id) // TODO: test for deletion in unsubscribed channel
	}
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

func (p *Publisher) Stopped() bool {
	return p.stopped
}

// func Save saves the state of the publisher to the configured saving location 
func (p *Publisher) Save() {
	//TODO
}