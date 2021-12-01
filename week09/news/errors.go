package news

// ErrArticleNotFound is returned when the article is not present in the publishing service.
type ErrArticleNotFound string

func (e ErrArticleNotFound) Error() string {
	return string(e)
}