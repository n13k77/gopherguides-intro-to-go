package news

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
)

// type Publisher publishes news articles to which can be subscribed
type Publisher struct {
	config 		PublisherConfig
	mutex  		sync.RWMutex
	subs   		map[string]map[int]chan Article
	src   		[]chan Article
	stopped		bool
	articles	map[int]Article
}

type PublisherState struct {
	Config 		PublisherConfig
	Categories 	[]string
	Stopped 	bool
	Articles	map[int]Article
}

type PublisherConfig struct {
	Backupfile 	string
	Publishfile string
	Jsonformat  bool
}


// func NewPublisher() creates a new instance of a Publisher
func NewPublisher(config PublisherConfig) (*Publisher) {
	// TODO error handling
	p := &Publisher{}

	p.subs = make(map[string]map[int]chan Article)
	p.stopped = false
	p.articles = make(map[int]Article)

	// if not specified, set backupfile to a default location
	if config.Backupfile == "" {
		config.Backupfile = "./backup.tmp"
	}
    
	p.config = config
	
	return p
}

// func Subscribe() adds a subscriber to a publisher. The subscriber has to
// provide the topic to which it is subscribing and its unique identifier
func (p *Publisher) Subscribe(id int, category string) (chan Article, error) {
	// TODO error handling
	p.mutex.Lock()
	defer p.mutex.Unlock()

	ch := make(chan Article)
	// use lowercase for category, avoid that SpoRtS and sports are different categories
	lc := strings.ToLower(category)	

	if _, exists := p.subs[lc][id]; exists {
		// already subscribed, silently ignore. 
		return nil, nil 
	}

	if _, exists := p.subs[lc]; exists {
		// The category exists for this publisher
		p.subs[lc][id] = ch
	} else {
		// The category does not exist for this publisher, first initialize map
		fmt.Println("category does not exists")
		p.subs[lc] = make(map[int]chan Article)
		p.subs[lc][id] = ch
	}
	return ch, nil
}

// func Unsubscriber() removes a subscriber from a publisher. The subscriber has to
// provide the topic from which it is unsubscribing and its own unique identifier
func (p *Publisher) Unsubscribe(id int, categories ...string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	cats := categories

	// if no categories are provided, remove all subscriptions for this subscriber
	if len(cats) == 0 {
		cats = p.Categories()
	}

	// TODO: how does the deletion from an unsubscribed channel work? Should NOT throw an error
	for _, cat := range cats {
		close(p.subs[cat][id])
		delete(p.subs[cat], id) 
	}
}

// func AddSource() adds a news source to the publisher
// TODO: as soon as a source publishes an article that has a previously unseen topic
// that topic should be added to the news publisher. 
func (p *Publisher) AddSource(s Source) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.src = append(p.src, s.GetSourceChannel())
}

// func Publish publishes an article to the subscribers subscribed to that particular topic
func (p *Publisher) Publish(a Article) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	if p.stopped {
		return
	}

	for subs := range p.subs[a.category] {
		p.subs[a.category][subs] <- a
	}

	// add it to the publisher archive
	p.articles[a.id] = a
}

// func Stop stops the publisher 
func (p *Publisher) Stop() {
	p.mutex.Lock()
	
	if !p.stopped {
		p.stopped = true
		for _, subs := range p.subs {
			for _, ch := range subs {
				// TODO: cancel subscribers as well
				close(ch)
			}
		}
	}
	
	p.mutex.Unlock()
	p.Save()
}

func (p *Publisher) Stopped() bool {
	return p.stopped
}

// func Save saves the state of the publisher to the configured saving location 
func (p *Publisher) Save() error {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	
	pc := PublisherState {
		Config: p.config,
		Articles: p.articles,
	}

	// fill the array with topics, save the topics since they are part of the state
	pc.Categories = p.Categories()

	// convert the config to json
	data, err := json.Marshal(pc)

	if err != nil {
		return err
	}
	
	err = os.WriteFile(pc.Config.Backupfile, data, 0644);

	if err != nil {
		return err
	}

	return nil
}

// func Categories returns all categories for which the publisher is publishing news
func (p *Publisher) Categories () []string {

	a := make([]string, 0, len(p.subs))
	
	for k := range p.subs {
		a = append(a, k)
	}

	return a
}

