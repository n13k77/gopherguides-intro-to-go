package news

// import (
// 	"encoding/json"
// 	"testing"
// )

// func TestArticleMarshallJson(t *testing.T) {
// 	tc := struct {
// 		name     string
// 		id       int
// 		title    string
// 		category string
// 		content  string
// 		result 	 string
// 	}{
// 		name:     "test article struct",
// 		id:       12345,
// 		title:    "test title",
// 		category: "test category",
// 		content:  "test content",
// 		result:   "{\"Id\":12345,\"Category\":\"test category\",\"Content\":\"test content\",\"Title\":\"test title\"}",
// 	}
// 	t.Run(tc.name, func(t *testing.T) {
// 		a := Article{tc.id, tc.category, tc.content, tc.title}
// 		u, err := json.Marshal(a)
// 		if err != nil {
// 			t.Fatalf("unexpected error occurred, got %s", err)
// 		}
		
// 		act := string(u)
// 		exp := tc.result

// 		if exp != act {
// 			t.Fatalf("expected %s, got %s", exp, act)
// 		}
// 	})
// }
