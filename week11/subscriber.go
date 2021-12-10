package news

import (
	"context"
	"log"
	"strings"
	"time"
)

type Subscriber struct {
	cats 	[]string
	id 		int
}

func NewSubscriber(categories []string) *Subscriber {
	log.Println("subscriber newsubscriber")

	// normalize categories

	lcs := []string{}
	for _, cat := range categories {
		lcs = append(lcs, strings.ToLower(cat))
	}

	s := &Subscriber{
		cats:	lcs,
		id:		int(time.Now().UnixNano()),
	}
	return s
}

func (s *Subscriber) Receive (p *Publisher) {
	log.Println("subscriber receive")

	ctx := context.Background()
	ctx, _ = context.WithCancel(ctx)
	ch, _ := p.Subscribe(ctx, s.cats)

	go func(){
		for a := range ch {
			log.Println("receive article")
			log.Println(a.String())
		}
	}()
}