package news

import (
	"testing"
)

func TestArticle(t *testing.T) {
	tc := struct {
		name     string
		id       int
		title    string
		category string
		content  string
	}{
		name:     "test article struct",
		id:       12345,
		title:    "test title",
		category: "test category",
		content:  "test content",
	}
	t.Run(tc.name, func(t *testing.T) {
		a := Article{tc.id, tc.category, tc.content, tc.title}

		if a.Category() != tc.category {
			t.Fatalf("expected category %s, got %s", tc.category, a.Category())
		}

		if a.Title() != tc.title {
			t.Fatalf("expected title %s, got %s", tc.title, a.Title())
		}

		if a.ID() != tc.id {
			t.Fatalf("expected ID %d, got %d", tc.id, a.ID())
		}

		if a.Content() != tc.content {
			t.Fatalf("expected content %s, got %s", tc.content, a.Content())
		}
	})
}

// TODO: Implement test function for article JSON marshalling / unmarshalling
// func TestArticleMarshallJson(t *testing.T) {
// 	testCases := []struct {
// 		desc	string
// 	}{
// 		{
// 			desc: "test json marshalling for article",
// 		},
// 	}
// 	for _, tC := range testCases {
// 		t.Run(tC.desc, func(t *testing.T) {

// 		})
// 	}
// }
