package news

import (
	"fmt"
	"testing"
)


func Test(t *testing.T) {
	testCases := []struct {
		desc	string
		config 	*PublisherConfig
		err 	error
	}{
		{desc: "save publisher, correct path", config: createConfig(t, "./test.txt", "./test.out"), err: nil},

	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			
			p := NewPublisher(*tc.config)
			s := NewSubscriber()

			s.Subscribe(p, "Horses")
			s.Subscribe(p, "Horses")
			s.Subscribe(p, "Horses2")
			
			fmt.Println(p.Categories())
			fmt.Println(s.Subscriptions())
			s.Unsubscribe(p)

			p.Stop()
		})
	}
}