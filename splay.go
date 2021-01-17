package splay

type Node struct {
}

type List struct {
}

func New() *List {
	return &List{}
}

func (l *List) Get(k uint64) interface{} {
	return nil
}

func (l *List) Set(k uint64, v interface{}) {
}

func (l *List) Has(k uint64) bool {
	return false
}

func (l *List) Del(k uint64) bool {
	return false
}
