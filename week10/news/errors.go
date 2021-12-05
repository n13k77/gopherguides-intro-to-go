package news

// ErrArticleNotFound is returned when the article is not present in the publishing service.
type ErrArticleNotFound int

func (e ErrArticleNotFound) Error() int {
	return int(e)
}

type ErrAlreadySubscribed string

func (e ErrAlreadySubscribed) Error () string {
	return string(e)
}