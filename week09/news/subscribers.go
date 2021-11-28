package news

import (
	"time"
)

type Subscriber struct {
	id 		int
}

func NewSubscriber () {
	s := &Subscriber{}
	s.id = int(time.Now().UnixMicro()) // unique ID for subscriber
}