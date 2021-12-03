package news

import (
	"fmt"
	"strings"
	"time"
)

type Subscriber struct {
	id 		int
	subs 	map[string]chan Article 	// save the combination of topic and channel
}

func NewSubscriber () *Subscriber {
	s := &Subscriber{
		id:		int(time.Now().UnixMicro()), // unique ID for subscriber
		subs:	make(map[string]chan Article),
	}

	for t, ch := range s.subs {
		// fmt.Println(t)
		go func(k string, c chan Article) {
			fmt.Printf("start receiving for %s\n", k)
		  	for art := range c {
				fmt.Println(art.String())
		 	}
		}(t, ch)
	}
	return s
}

func (s *Subscriber) Subscribe (p *Publisher, categories ...string) error {
	// TODO error handling
	for _, category := range categories {
		ch, err := p.Subscribe(s.id, category)
		if err != nil {
			return err
		}
		s.subs[strings.ToLower(category)] = ch
	}
	return nil
}

func (s *Subscriber) Unsubscribe (p *Publisher, categories ...string) error {
	// TODO error handling
	p.Unsubscribe(s.id, categories...)
	return nil
}

// func (s *Subscriber) Receive() {
// 	for t, ch := range s.subs {
// 		// fmt.Println(t)
// 		go func(k string, c chan Article) {
// 			fmt.Printf("start receiving for %s\n", k)
// 		  	for art := range c {
// 				fmt.Println(art.String())
// 		 	}
// 		}(t, ch)
// 	}
// }

func (s *Subscriber) Subscriptions () []string {

	a := make([]string, 0, len(s.subs))

	for k := range s.subs {
		a = append(a, k)
	}

	return a
}
