package news

import (
	"fmt"
)

// ErrArticleNotFound is returned when the article is not present in the publishing service.
type ErrArticleNotFound int

func (e ErrArticleNotFound) Error() string {
	return fmt.Sprintf("article %d not found", e)
}

// ErrAlreadySubscribed is returned when the subscription already exists.
type ErrAlreadySubscribed string

func (e ErrAlreadySubscribed) Error () string {
	return string(e)
}