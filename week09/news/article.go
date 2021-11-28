package news

import (
	"encoding/json"
	"fmt"
)

type Article struct {
	id			int
	category 	string
	content 	string
	title		string
}

func (a Article) Title() string {
	return a.title
}

func (a Article) ID() int {
	return a.id
}

func (a Article) Category() string {
	return a.category
}

func (a Article) Content() string {
	return a.content
}

func (a Article) String() string {
	return fmt.Sprintf("id: %d, title: %s, category: %s, content: %s", a.id, a.category, a.content, a.title)
}

func (a Article) MarshallJSON() ([]byte, error)  {
    j, err := json.Marshal(struct {
		Id			int
		Category 	string
		Content 	string
		Title		string
	}{
		Id:			a.id,
		Category: 	a.category,
		Content: 	a.content,
		Title:		a.title,
    })
    if err != nil {
           return nil, err
    }
    return j, nil
}