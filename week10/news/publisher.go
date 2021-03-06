package news

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

// type Publisher publishes news articles to which can be subscribed
type Publisher struct {
	config 		PublisherConfig
	mutex  		sync.RWMutex
	subs   		map[string]map[int]chan Article
	src   		[]<-chan Article
	stopped		bool
	articles	[]Article
}

type PublisherState struct {
	Config 		PublisherConfig
	Categories 	[]string
	Stopped 	bool
	Articles	[]Article
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

	p.subs = make(map[string]map[int]chan Article)
	p.stopped = false

	// if not specified, set backupfile to a default location
	if config.Backupfile == "" {
		config.Backupfile = "./backup.tmp"
	}
    
	p.config = config

	return p
}

// func Start starts the publisher. If no backup file is provided, a new 
// instance is created of the Publisher type. If a backup file is provided,
// settings from that file are used

func Start(file string) (*Publisher, error){
	log.Println("publisher start")

	if file == "" {
		log.Println("publisher start publisher new config")
		p := NewPublisher(PublisherConfig{})
		return p, nil
	}

	log.Println("publisher start publisher existing config")
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var ps PublisherState
	json.Unmarshal([]byte(byteValue),&ps)

	p := NewPublisher(ps.Config)
	p.articles = ps.Articles

	log.Println("bla")
	return p, nil
}

// func Stop stops the publisher 
func (p *Publisher) Stop() {
	log.Println("publisher stop")
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

// func Stopped returns whether the publisher stopped or not
func (p *Publisher) Stopped() bool {
	log.Println("publisher stopped")
	return p.stopped
}

// func Save saves the state of the publisher to the configured saving location 
func (p *Publisher) Save() error {
	log.Println("publisher save")

	p.mutex.RLock()
	defer p.mutex.RUnlock()

	var ps PublisherState

	ps.Articles = p.articles
	ps.Categories = p.Categories()
	ps.Config = p.config

	// ps := PublisherState {
	// 	Config: p.config,
	// 	Articles: p.articles,
	// }

	// convert the config to json
	data, err := json.Marshal(ps)

	if err != nil {
		return err
	}
	

	err = os.WriteFile(ps.Config.Backupfile, data, 0644);

	if err != nil {
		return err
	}

	return nil
}


// func Subscribe() adds a subscriber to a publisher. The subscriber has to
// provide the topic to which it is subscribing and its unique identifier
func (p *Publisher) Subscribe(id int, category string) (<-chan Article, error) {
	log.Println("publisher subscribe")
	// this is quite a broad lock, a lock with smaller scope did not function
	p.mutex.Lock()
	defer p.mutex.Unlock()

	ch := make(chan Article)
	// use lowercase for category, guarantees that SpoRtS and sports are same categories
	lc := strings.ToLower(category)	

	if _, exists := p.subs[lc][id]; exists {
		// already subscribed, return error. 
		return nil, ErrAlreadySubscribed(lc) 
	}

	if _, exists := p.subs[lc]; exists {
		// The category exists for this publisher
		p.subs[lc][id] = ch
	} else {
		// The category does not exist for this publisher, first initialize map
		log.Println("category does not exist, create category")
		p.subs[lc] = make(map[int]chan Article)
		log.Println("category does not exist, create subscription")
		p.subs[lc][id] = ch
	}
	return ch, nil
}

// func Unsubscribe() removes a subscriber from a publisher. The subscriber has to
// provide its own unique identifier and the categories from which it is unsubscribing.
// If no categories are provided, all subscriptions will be removed, and the channels 
// closed.
// If a subscriber wants to unsubscribe from a channel it was not subscribed to,
// no explicit error will be thrown because the end situation matches the desired
// situation.
func (p *Publisher) Unsubscribe(id int, categories ...string) {
	log.Printf("publisher unsubscribe %d\n", id)

	cats := categories

	// if no categories are provided, remove all subscriptions for this subscriber
	if len(cats) == 0 {
		cats = p.Categories()
	}

	for _, cat := range cats {
		log.Printf("publisher unsubscribe category %s", cat)

		lc := strings.ToLower(cat)
		
		// remove subscriber from category if it was subcribed
		if _, ok := p.subs[lc][id]; ok {
			log.Printf("publisher remove subscriber %d from category %s", id, lc)
			p.mutex.Lock()
			// close subscriber channel
			close(p.subs[cat][id])
			// remove subscriber from category map
			delete(p.subs[cat], id) 
			p.mutex.Unlock()
		} 
	}
}

// func AddSource() adds a news source to the publisher, distributes its articles
func (p *Publisher) DistributeSource(s Source) {
	log.Println("publisher distribute source")

	p.mutex.Lock()
	p.src = append(p.src, s.GetSourceChannel())
	p.mutex.Unlock()

	go func() {
		// start listening to articles that are published by the source
		log.Println("publisher distribute source start listening")
		for a := range s.GetSourceChannel() {

			// add the received article to the archive 
			p.mutex.Lock()
			a.id = len(p.articles)// + 1
			p.articles = append(p.articles, a)
			p.mutex.Unlock()

			lc := strings.ToLower(a.Category())
			if _, ok := p.subs[lc]; ok {
				log.Println("publisher distribute source publish existing category")
				p.mutex.RLock()
				for _, ch := range p.subs[lc] {
					ch<- a
				}
				p.mutex.RUnlock()
			} else {
				log.Println("publisher distribute source publish new category")
				p.mutex.Lock()
				p.subs[lc] = make(map[int]chan Article)
				p.mutex.Unlock()
			}
		}
	}()
}

// func Categories returns all categories for which the publisher is publishing news
func (p *Publisher) Categories () []string {
	log.Println("publisher categories")

	a := make([]string, 0, len(p.subs))
	
	for k := range p.subs {
		a = append(a, k)
	}
	
	return a
}

// func Articles returns the articles from the archive as specified by the IDs
func (p *Publisher) Articles (ids ...int) []Article {
	log.Printf("publisher articles")

	p.mutex.RLock()
	defer p.mutex.RUnlock()
	a := []Article{}

	for _, id := range ids {
		log.Printf("publisher articles id %d", id)
		a = append(a, p.articles[id])	
	}
	return a
}

// func Clear clears the publisher archive
func (p *Publisher) Clear () {
	log.Printf("publisher clear")
	p.mutex.Lock()
	defer p.mutex.Unlock()
	
	p.articles = nil
}
