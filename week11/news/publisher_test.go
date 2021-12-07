package news

import (
	"testing"
)

func createConfig(t testing.TB, backupfile string, publishfile string) *PublisherConfig {
	t.Helper()
	return &PublisherConfig{
		Jsonformat: 	false,
		Backupfile: 	backupfile,
		Publishfile: 	publishfile,
	}
}