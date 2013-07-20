package domain

type Ranker interface {
	Rank(Document) float32
}
