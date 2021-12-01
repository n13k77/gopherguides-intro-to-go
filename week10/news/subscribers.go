package news

import (
	"time"
)

type Subscriber struct {
	id 		int
	subs 	map[string]<-chan Article 	// save the combination of topic and channel
}

func NewSubscriber () *Subscriber {
	return &Subscriber{
		id:		int(time.Now().UnixMicro()), // unique ID for subscriber
		subs:	make(map[string]<-chan Article),
	}
}

func (s *Subscriber) Subscribe (p *Publisher, categories ...string) error {
	// TODO error handling
	for _, category := range categories {
		ch, err := p.Subscribe(s.id, category)
		if err != nil {
			return err
		}
		s.subs[category] = ch
	}
	return nil
}

func (s *Subscriber) Unsubscribe (p *Publisher, categories ...string) error {
	// TODO error handling
	p.Unsubscribe(s.id, categories...)
	return nil
}

func (s *Subscriber) Subscriptions () []string {

	a := make([]string, 0, len(s.subs))
	
	for k, _ := range s.subs {
		a = append(a, k)
	}

	return a
}
