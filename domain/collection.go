package domain

type Collection interface {
	Search(query string, ranker Ranker) (Document, error)
	Find(DocumentId) (Document, error)
}
