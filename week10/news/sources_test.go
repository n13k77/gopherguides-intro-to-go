package news

import (
	"testing"
)

func TestRandomSourcePublish(t *testing.T) {
	tc := struct {
		desc	string
	}{
		desc: "test random publishing source",
	}
	t.Run(tc.desc, func(t *testing.T) {
		rs := &RandomSource{}
		art := []Article{}

		go func() {
			for rcv := range rs.ch {
				art = append(art, rcv)
				if len(art) == 6 {
					rs.Stop()
				}
			}
		}()
		rs.Publish()
	})
}