package news

import (
	"encoding/json"
	"fmt"
)

type Article struct {
	Id			int
	Category 	string
	Content 	string
	Title		string
}

func (a Article) String() string {
	return fmt.Sprintf("id: %d, title: %s, category: %s, content: %s", a.Id, a.Title, a.Category, a.Content)
}

func (a Article) MarshalJSON() ([]byte, error) {
    arr := []interface{}{a.Id, a.Title, a.Category, a.Content}
    return json.Marshal(arr)
}