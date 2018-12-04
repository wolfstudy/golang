package skiplist

type Lesser interface {
	Less(other Lesser) bool
}

type MapContainer interface {
	Search(Lesser) (Lesser, bool)
	ReplaceOrInsert(Lesser) Lesser
	Delete(Lesser) (deleted Lesser, found bool)
	Min() (less Lesser, found bool)
	Max() (less Lesser, found bool)
	Ascend(iterHandler func(i Lesser) bool)
	Len() int
	DeleteMax() (less Lesser, found bool)
}
