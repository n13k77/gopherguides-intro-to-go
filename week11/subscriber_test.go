package news

import (
	"context"
	"testing"
	"time"
)

func TestSubscriberRandomSource(t *testing.T) {
	tc := struct {
		desc	string
		config 	*PublisherConfig
	}{
		desc: "test article receive for subscriber", 
		config: createConfig(t, "./test.txt", "./test.out"),
	}
	t.Run(tc.desc, func(t *testing.T) {

		ctx := context.Background()
		ctx, _ = context.WithCancel(ctx)

		
		p := NewPublisher(*tc.config)
		defer p.Stop()

		sub1 := NewSubscriber([]string{"World", "ecoNomics"})
		sub2 := NewSubscriber([]string{"cookinG", "ecOnomics"})
		src := NewRandomSource()
		
		p.Dispatch(ctx, src)
		defer p.cancel()
		
		go func(){
			sub1.Receive(p)
			sub2.Receive(p)
		}()

		time.Sleep(time.Second * 1)
	})
}

func TestSubscriberFileSource(t *testing.T) {
	tc := struct {
		desc	string
		config 	*PublisherConfig
	}{
		desc: "test article receive for subscriber", 
		config: createConfig(t, "./test.txt", "./test.out"),
	}
	t.Run(tc.desc, func(t *testing.T) {

		ctx := context.Background()
		ctx, _ = context.WithCancel(ctx)

		p := NewPublisher(*tc.config)
		sub1 := NewSubscriber([]string{"World", "ecoNomics"})
		sub2 := NewSubscriber([]string{"cookinG", "ecOnomics"})
		src := NewFileSource("./tmp")
		
		p.Dispatch(ctx, src)
		
		go func(){
			sub1.Receive(p)
			sub2.Receive(p)
		}()

		time.Sleep(time.Second * 10)

	})
}
