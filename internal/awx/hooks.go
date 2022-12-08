package awx

type Callee uint8
type Source uint8

const (
	CalleeCreate Callee = iota
	CalleeUpdate
	CalleeRead
	CalleeDelete
)

const (
	SourceData Source = iota
	SourceResource
)
