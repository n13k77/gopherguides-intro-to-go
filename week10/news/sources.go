package news

import (
	"fmt"
	"time"
)

type Source interface {
	GetSourceChannel() (chan Article)
	Publish() error
	Stop() error
}

type RandomSource struct {
	articles 	[]Article
	ch 			chan Article
}

func NewRandomSource() *RandomSource {
	return &RandomSource {
		ch: make(chan Article),
		articles: []Article{
			{id: 1, category: "world", title: "Ethiopia's Tigray conflict: Lalibela retaken", content: "Ethiopian troops have recaptured the historic town of Lalibela from Tigrayan rebels, the government has said."},
			{id: 2, category: "sports", title: "Max Verstappen & Lewis Hamilton set for thrilling Formula 1 finale", content: "The most intense Formula 1 championship fight for years will be decided over the next two weekends in the Middle East."},
			{id: 3, category: "local", title: "Clear Flour Bread serves the best cookies in Mass., according to Yelp", content: "It’s officially holiday cookie season, that time of year when Christmas tree-shaped sugar cookies lay snugly in tins next to snickerdoodles and gingersnap cookies."},
			{id: 4, category: "cooking", title: "Salmon and broccoli pasta", content: "A simple salmon pasta that’s ready in under 15 minutes. This recipe makes two generous servings or three lighter meals. It’s also very easy to double up."},
			{id: 5, category: "economics", title: "Tel Aviv named as world's most expensive city to live in", content: "Tel Aviv has been named as the most expensive city in the world to live in, as soaring inflation and supply-chain problems push up prices globally."},
			{id: 6, category: "world", title: "Rust: US Police to search arms supplier over fatal film shooting", content: "Police investigating the fatal shooting on the set of the Alec Baldwin movie Rust have obtained a further warrant to search the premises of an arms supplier in the US."},
			
		},
	}
}

func (rs *RandomSource) GetSourceChannel() chan Article {
	return rs.ch
}

func (rs *RandomSource) Publish() error {
	for _, a := range rs.articles {
		fmt.Println(a.String())
		rs.ch <- a
		time.Sleep(1 * time.Second)
	}
	return nil 
}

func (rs *RandomSource) Stop() error {
	close(rs.ch)
	return nil
}