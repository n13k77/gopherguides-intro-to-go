package news

import (
	//"errors"
	"os"
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

func TestSavePublisher(t *testing.T) {
	testCases := []struct {
		desc	string
		config  *PublisherConfig
		err 	error
	}{
		{desc: "save publisher, correct path", config: createConfig(t, "./test.txt", "./test.out"), err: nil},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			p := NewPublisher(*tc.config)

			p.Save()

			act, err := os.ReadFile(tc.config.Backupfile)
			if err != nil {
				t.Fatal(err)
			}

			exp := "{\"Config\":{\"Backupfile\":\"./test.txt\",\"Publishfile\":\"./test.out\",\"Jsonformat\":false},\"Categories\":[],\"Stopped\":false,\"Articles\":{}}"
			if exp != string(act) {
				t.Fatalf("expected %s, got %s", exp, act)
			}

			// if a Publisher is created during the test run, clean it up
			if p != nil && ! p.Stopped() {
				p.Stop()
			}
		})
	}
}

