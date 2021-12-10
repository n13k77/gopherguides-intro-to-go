package news

import (
	"context"
	"log"
)


type RandomSource struct {
	articles 	[]Article
	ch 			chan Article
}

func NewRandomSource() *RandomSource {
	log.Println("sources newrandomsource")
	return &RandomSource {
		ch: make(chan Article),
		articles: []Article{
			{Category: "world", Title: "Ethiopia's Tigray conflict: Lalibela retaken", Content: "Ethiopian troops have recaptured the historic town of Lalibela from Tigrayan rebels, the government has said."},
			{Category: "sports", Title: "Max Verstappen & Lewis Hamilton set for thrilling Formula 1 finale", Content: "The most intense Formula 1 championship fight for years will be decided over the next two weekends in the Middle East."},
			{Category: "local", Title: "Clear Flour Bread serves the best cookies in Mass., according to Yelp", Content: "It’s officially holiday cookie season, that time of year when Christmas tree-shaped sugar cookies lay snugly in tins next to snickerdoodles and gingersnap cookies."},
			{Category: "cooking", Title: "Salmon and broccoli pasta", Content: "A simple salmon pasta that’s ready in under 15 minutes. This recipe makes two generous servings or three lighter meals. It’s also very easy to double up."},
			{Category: "economics", Title: "Tel Aviv named as world's most expensive city to live in", Content: "Tel Aviv has been named as the most expensive city in the world to live in, as soaring inflation and supply-chain problems push up prices globally."},
			{Category: "world", Title: "Rust: US Police to search arms supplier over fatal film shooting", Content: "Police investigating the fatal shooting on the set of the Alec Baldwin movie Rust have obtained a further warrant to search the premises of an arms supplier in the US."},
		},
	}
}

func (rs *RandomSource) ConnectSource(ctx context.Context) <-chan Article {
	log.Println("sources connectsource")

	go func(ctx context.Context){
		i := 0
		for {
			select {
			case <-ctx.Done():
				rs.Stop()
			case rs.ch <- rs.articles[i]:
				log.Printf("sources publish article %s", rs.articles[i].Title)
				i = (i + 1) % 5 
			}
		}
	}(ctx)

	return rs.ch
}

func (rs *RandomSource) Stop() {
	log.Println("sources stop")
	close(rs.ch)
}