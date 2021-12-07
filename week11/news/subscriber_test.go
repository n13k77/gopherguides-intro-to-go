package news

import (
	"testing"
)

func TestSubscriberArticleReceive(t *testing.T) {
	tc := struct {
		desc	string
		config 	*PublisherConfig
	}{
		desc: "test article receive for subscriber", 
		config: createConfig(t, "./test.txt", "./test.out"),
	}
	t.Run(tc.desc, func(t *testing.T) {

		p := NewPublisher(*tc.config)
		sub1 := NewSubscriber([]string{"World", "ecoNomics"})
		sub2 := NewSubscriber([]string{"cookinG", "ecOnomics"})
		src := NewRandomSource()
		
		defer func() {
			// sub1.Unsubscribe(p)
			// sub2.Unsubscribe(p)
			// src.Stop()
			//p.Stop()
		}()
		
		go func() {
			src.Publish()
			p.Dispatch(src)
		}()			
		go func(){ 
			sub1.Receive(*p)
			sub2.Receive(*p)
		}()
		


	})
}