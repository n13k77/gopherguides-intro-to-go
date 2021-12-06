package news

import (
	"fmt"
	"log"
	"time"
)

type Subscriber struct {
	id 		int
	cats 	[]string
}

func NewSubscriber(categories ...string) *Subscriber {
	log.Println("subscriber newsubscriber")
	s := &Subscriber{
		id:		int(time.Now().UnixNano()), // unique ID for subscriber
		cats:	categories,
	}
	return s
}

func (s *Subscriber) Receive (p *Publisher) {
	ch, _ := p.Subscribe(s.cats)

	for a := range ch {
		fmt.Println(a.String())
	}
}