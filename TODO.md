

Server Command
- routes
-- GET    /{collection}/search/{term}
-- DELETE /{collection}/documents/{id}
-- GET    /{collection}/documents/{id}
-- PUT    /{collection}/documents/{id}

- http listen and serve

Result struct
	Id DocumentId
	Weight int 

Collection Interface
	- Search(string, Ranker?) (Results, error)
	- All(DocumentIds) (Documents, error)
	- Find(DocumentId) (Document, error)
	- Destroy(DocumentId) error

Index Interfaces

-- InvertedIndex
-- ForwardIndex
