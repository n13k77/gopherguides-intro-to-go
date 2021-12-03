package news

import (
	//"fmt"
	"sort"
	"testing"
)


func TestSubscriber(t *testing.T) {
	tc := struct {
		desc	string
		config 	*PublisherConfig
	}{
		desc: "test subscription", 
		config: createConfig(t, "./test.txt", "./test.out"),
	}
	t.Run(tc.desc, func(t *testing.T) {
		
		p := NewPublisher(*tc.config)
		s := NewSubscriber()

		defer func() {
			s.Unsubscribe(p)
			p.Stop()
		}()

		s.Subscribe(p, "World")
		s.Subscribe(p, "EcoNoMicS")
		s.Subscribe(p, "economics")
		
		act_sub := s.Subscriptions()
		act_pub := p.Categories()
		exp := []string{"world","economics"}

		sort.Strings(exp)
		sort.Strings(act_sub)
		sort.Strings(act_pub)

		// verify the categories in the subscriber
		if len(act_sub) != len(exp) {
			t.Fatalf("categories for subscriber do not match, expected %s, got %s", exp, act_sub)
		}
		for i, v := range act_sub {
			if v != exp[i] {
				t.Fatalf("categories for subscriber do not match, expected %s, got %s", exp, act_sub)
			}
		}

		// verify the categories in the publisher
		if len(act_pub) != len(exp) {
			t.Fatalf("categories for publisher do not match, expected %s, got %s", exp, act_pub)
		}
		for i, v := range act_pub {
			if v != exp[i] {
				t.Fatalf("categories for publisher do not match, expected %s, got %s", exp, act_pub)
			}
		}
	})
}

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
		sub := NewSubscriber()
		src := NewRandomSource()
		
		p.AddSource(src)

		defer func() {
			sub.Unsubscribe(p)
			src.Stop()
			p.Stop()
		}()

		sub.Subscribe(p, "World")
		sub.Subscribe(p, "Economics")

		src.Publish()
	})
}