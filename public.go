package collection

import (
	"golang.org/x/exp/constraints"

	collection "github.com/idevelopthings/collection/keyed_collection"
)

func (col *Collection[T]) Add(item T) *Collection[T] {
	col.Lock()
	defer col.Unlock()

	return col.add(item)
}

func (col *Collection[T]) Remove(item T) *Collection[T] {
	col.Lock()
	defer col.Unlock()

	return col.remove(item)
}

func (col *Collection[T]) Contains(item T) bool {
	col.Lock()
	defer col.Unlock()

	return col.contains(item)
}

func (col *Collection[T]) Items() []*T {
	col.Lock()
	defer col.Unlock()

	return col.items
}

func (col *Collection[T]) Len() int {
	col.Lock()
	defer col.Unlock()

	return len(col.items)
}

func (col *Collection[T]) IsEmpty() bool {
	col.Lock()
	defer col.Unlock()

	return len(col.items) == 0
}

func (col *Collection[T]) IsNotEmpty() bool {
	col.Lock()
	defer col.Unlock()

	return len(col.items) > 0
}

func (col *Collection[T]) First() *T {
	col.Lock()
	defer col.Unlock()

	return col.items[0]
}

func (col *Collection[T]) Last() *T {
	col.Lock()
	defer col.Unlock()

	return col.items[len(col.items)-1]
}

func (col *Collection[T]) AddRange(items []T) *Collection[T] {
	col.Lock()
	defer col.Unlock()

	for _, item := range items {
		col.add(item)
	}

	return col
}

func (col *Collection[T]) Clear() *Collection[T] {
	col.Lock()
	defer col.Unlock()

	col.items = make([]*T, 0)

	return col
}

func (col *Collection[T]) Where(predicate func(item T) bool) []*T {
	col.Lock()
	defer col.Unlock()

	return col.where(predicate)
}

func (col *Collection[T]) Count(predicate func(item T) bool) int {
	col.Lock()
	defer col.Unlock()

	return col.count(predicate)
}

func (col *Collection[T]) Any(predicate func(item T) bool) bool {
	col.Lock()
	defer col.Unlock()

	return col.any(predicate)
}

func (col *Collection[T]) All(predicate func(item T) bool) bool {
	col.Lock()
	defer col.Unlock()

	return col.all(predicate)
}

func (col *Collection[T]) FirstWhere(predicate func(item T) bool) *T {
	col.Lock()
	defer col.Unlock()

	return col.firstWhere(predicate)
}

func (col *Collection[T]) LastWhere(predicate func(item T) bool) *T {
	col.Lock()
	defer col.Unlock()

	return col.lastWhere(predicate)
}

func (col *Collection[T]) Map(mapFunc func(index int, item T) T) *Collection[T] {
	col.Lock()
	defer col.Unlock()

	return col.mapFunc(mapFunc)
}

func (col *Collection[T]) ForEach(forEachFunc func(index int, item T)) *Collection[T] {
	col.Lock()
	defer col.Unlock()

	return col.forEach(forEachFunc)
}

func KeyBy[Key constraints.Ordered, Val comparable](col *Collection[Val], mapFunc func(index int, item Val) Key) *collection.KeyedCollection[Key, Val] {
	col.Lock()
	defer col.Unlock()

	entries := make(map[Key]Val, 0)
	for i, item := range col.items {
		keyResult := mapFunc(i, *item)

		entries[keyResult] = *item
	}

	return collection.New[Key, Val](entries)
}

func (col *Collection[T]) Set(index int, value T) {
	col.Lock()
	defer col.Unlock()

	col.items[index] = &value
}

func (col *Collection[T]) IndexOf(item T) int {
	col.Lock()
	defer col.Unlock()

	for i, v := range col.items {
		if *v == item {
			return i
		}
	}

	return -1
}

func (col *Collection[T]) Filter(filterFunc func(index int, item T) bool) []T {
	col.Lock()
	defer col.Unlock()

	items := make([]T, 0)

	for i, v := range col.items {
		if filterFunc(i, *v) {
			items = append(items, *v)
		}
	}

	return items
}
