package domain

type DocumentId string

type Document struct {
	Id    DocumentId
	Terms []string
}
