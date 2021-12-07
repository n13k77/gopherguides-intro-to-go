package news

import (
	"fmt"
	"log"
)

type Subscriber struct {
	cats 	[]string
}

func NewSubscriber(categories ...string) *Subscriber {
	log.Println("subscriber newsubscriber")
	s := &Subscriber{
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