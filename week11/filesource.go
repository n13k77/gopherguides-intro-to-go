package news

import (
	"context"
	"log"
	// "os"
	"time"
)

type FileSource struct {
	path 	string
	ch 		chan Article
}

func NewFileSource(path string) *FileSource {
	log.Println("sources newfilesource")
	return &FileSource{path: path}
}

func (fs *FileSource) ConnectSource(ctx context.Context) <-chan Article {
	log.Println("filesource connectsource")

	go func(ctx context.Context){
		for {
			select {
			case <-ctx.Done():
				fs.Stop()
			case fs.ch <- fs.WatchDir(fs.path):
				log.Printf("checking dir %s", fs.path)
			case <-time.After(1 * time.Second):
				log.Printf("waiting")
			}
		}
	}(ctx)

	return fs.ch
}

func (rs *FileSource) WatchDir (path string) Article {
	return Article{Id: 0, Category: "world", Title: "Ethiopia's Tigray conflict: Lalibela retaken", Content: "Ethiopian troops have recaptured the historic town of Lalibela from Tigrayan rebels, the government has said."}
}

func (rs *FileSource) Stop() {
	log.Println("sources stop")
	close(rs.ch)
}	