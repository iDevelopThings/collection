package collection

import (
	"reflect"
	"sync"
)

type Collection[T comparable] struct {
	*sync.Mutex

	items []*T

	itemType reflect.Type
}

func New[T comparable](items ...T) *Collection[T] {
	col := &Collection[T]{
		Mutex: &sync.Mutex{},
		items: make([]*T, 0),
	}

	if len(items) > 0 {
		col.AddRange(items)
	}

	col.itemType = reflect.TypeOf((*T)(nil)).Elem()

	return col
}
