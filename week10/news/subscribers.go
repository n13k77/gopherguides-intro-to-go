package news

import (
	"log"
	"strings"
	"time"
)

type Subscriber struct {
	id 		int
	subs 	map[string]chan Article 	// save the combination of topic and channel
}

func NewSubscriber () *Subscriber {
	log.Println("subscriber newsubscriber")
	s := &Subscriber{
		id:		int(time.Now().UnixNano()), // unique ID for subscriber
		subs:	make(map[string]chan Article),
	}
	return s
}

func (s *Subscriber) Subscribe (p *Publisher, categories ...string) error {
	for _, category := range categories {
		ch, err := p.Subscribe(s.id, category)
		if err != nil {
			return err
		}
		log.Printf("subscriber subscribe %s", strings.ToLower(category))
		s.subs[strings.ToLower(category)] = ch

		go func() {
			for a := range ch {
			log.Printf("subscriber subcribe recieved article %d", a.ID())
			}
		}()

	}
	return nil
}

func (s *Subscriber) Unsubscribe (p *Publisher, categories ...string) error {
	log.Println("subscriber unsubscribe")
	p.Unsubscribe(s.id, categories...)
	return nil
}

func (s *Subscriber) Subscriptions () []string {
	log.Println("subscriber subscriptions")
	a := make([]string, 0, len(s.subs))

	for k := range s.subs {
		a = append(a, k)
	}

	return a
}
