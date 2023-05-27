package collection

func (col *KeyedCollection[K, T]) GetByKey(key K) *T {
	col.Lock()
	defer col.Unlock()

	return col.getByKey(key)
}

func (col *KeyedCollection[K, T]) GetByItem(item *T) *T {
	col.Lock()
	defer col.Unlock()

	return col.getByItem(item)
}

func (col *KeyedCollection[K, T]) Add(key K, item T) *KeyedCollection[K, T] {
	col.Lock()
	defer col.Unlock()

	return col.add(key, item)
}

func (col *KeyedCollection[K, T]) RemoveByKey(key K) *KeyedCollection[K, T] {
	col.Lock()
	defer col.Unlock()

	return col.removeByKey(key)
}

func (col *KeyedCollection[K, T]) RemoveByItem(item T) *KeyedCollection[K, T] {
	col.Lock()
	defer col.Unlock()

	return col.removeByItem(item)
}

func (col *KeyedCollection[K, T]) ContainsKey(key K) bool {
	col.Lock()
	defer col.Unlock()

	return col.containsKey(key)
}

func (col *KeyedCollection[K, T]) ContainsItem(item T) bool {
	col.Lock()
	defer col.Unlock()

	return col.containsItem(item)
}

func (col *KeyedCollection[K, T]) Values() []T {
	col.Lock()
	defer col.Unlock()

	return col.values()
}

func (col *KeyedCollection[K, T]) Items() map[K]*T {
	col.Lock()
	defer col.Unlock()

	return col.items.Map()
}

func (col *KeyedCollection[K, T]) Len() int {
	col.Lock()
	defer col.Unlock()

	return col.items.Len()
}

func (col *KeyedCollection[K, T]) IsEmpty() bool {
	col.Lock()
	defer col.Unlock()

	return col.items.Len() == 0
}

func (col *KeyedCollection[K, T]) IsNotEmpty() bool {
	col.Lock()
	defer col.Unlock()

	return col.items.Len() > 0
}

func (col *KeyedCollection[K, T]) First() *T {
	col.Lock()
	defer col.Unlock()

	return col.first()
}

func (col *KeyedCollection[K, T]) Last() *T {
	col.Lock()
	defer col.Unlock()

	return col.last()
}

func (col *KeyedCollection[K, T]) AddRange(items map[K]T) *KeyedCollection[K, T] {
	col.Lock()
	defer col.Unlock()

	for key, item := range items {
		col.add(key, item)
	}

	return col
}

func (col *KeyedCollection[K, T]) Clear() *KeyedCollection[K, T] {
	col.Lock()
	defer col.Unlock()

	col.items = NewSortedMapOf[K, T]()

	return col
}

func (col *KeyedCollection[K, T]) Keys() []K {
	col.Lock()
	defer col.Unlock()

	keys := make([]K, 0)

	for _, kv := range col.items.entries {
		keys = append(keys, kv.Key)
	}

	return keys
}

func (col *KeyedCollection[K, T]) Where(predicate func(key K, item T) bool) *KeyedCollection[K, T] {
	col.Lock()
	defer col.Unlock()

	return col.where(predicate)
}

func (col *KeyedCollection[K, T]) Count(predicate func(key K, item T) bool) int {
	col.Lock()
	defer col.Unlock()

	return col.count(predicate)
}

func (col *KeyedCollection[K, T]) Any(predicate func(key K, item T) bool) bool {
	col.Lock()
	defer col.Unlock()

	return col.any(predicate)
}

func (col *KeyedCollection[K, T]) All(predicate func(key K, item T) bool) bool {
	col.Lock()
	defer col.Unlock()

	return col.all(predicate)
}

func (col *KeyedCollection[K, T]) FirstWhere(predicate func(key K, item T) bool) *T {
	col.Lock()
	defer col.Unlock()

	return col.firstWhere(predicate)
}

func (col *KeyedCollection[K, T]) LastWhere(predicate func(key K, item T) bool) *T {
	col.Lock()
	defer col.Unlock()

	return col.lastWhere(predicate)
}

func (col *KeyedCollection[K, T]) Map(mapFunc func(key K, item T) T) *KeyedCollection[K, T] {
	col.Lock()
	defer col.Unlock()

	return col.mapFunc(mapFunc)
}

func (col *KeyedCollection[K, T]) ForEach(forEachFunc func(key K, item T)) *KeyedCollection[K, T] {
	col.Lock()
	defer col.Unlock()

	return col.forEach(forEachFunc)
}
