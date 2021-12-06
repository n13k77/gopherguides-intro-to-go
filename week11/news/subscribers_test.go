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

func TestSubscriberDoubleSubscription(t *testing.T) {
	tc := struct {
		desc	string
		config 	*PublisherConfig
	}{
		desc: "test double subscription", 
		config: createConfig(t, "./test.txt", "./test.out"),
	}
	t.Run(tc.desc, func(t *testing.T) {
		
		p := NewPublisher(*tc.config)
		s := NewSubscriber()

		defer func() {
			s.Unsubscribe(p)
			p.Stop()
		}()

		err := s.Subscribe(p, "EcoNoMicS")

		if err != nil {
			t.Fatalf("unexpected error setting up test %s", tc.desc)
		}

		act := s.Subscribe(p, "economics")
		exp := ErrAlreadySubscribed("economics")

		if exp != act {
			t.Fatalf("expected %s, got %s", exp, act)
		}

	})
}

// for properly testing the subscriber receiver functionality, the whole chain of 
// source -> publisher -> subscriber is created here. So this ends up being more 
// of an integration test. 
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
		sub1 := NewSubscriber()
		sub2 := NewSubscriber()
		src := NewRandomSource()
		
		p.DistributeSource(src)
		
		defer func() {
			sub1.Unsubscribe(p)
			sub2.Unsubscribe(p)
			src.Stop()
			p.Stop()
		}()

		sub1.Subscribe(p, "World")
		sub1.Subscribe(p, "Economics")
		sub2.Subscribe(p, "cooking")
		src.Publish()
	})
}
