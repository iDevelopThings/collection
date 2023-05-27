package collection

func (col *KeyedCollection[K, T]) getByKey(key K) *T {
	return col.items.Get(key)
}

func (col *KeyedCollection[K, T]) getByItem(item *T) *T {
	for _, v := range col.items.values {
		if v == *item {
			return item
		}
	}

	return nil
}

func (col *KeyedCollection[K, T]) first() *T {
	for _, v := range col.items.values {
		return &v
	}

	return nil
}

func (col *KeyedCollection[K, T]) last() *T {
	values := col.items.Values()

	if len(values) == 0 {
		return nil
	}

	val := values[col.items.Len()-1]

	return &val
}

func (col *KeyedCollection[K, T]) add(key K, item T) *KeyedCollection[K, T] {
	col.items.Add(key, item)

	return col
}

func (col *KeyedCollection[K, T]) removeByKey(key K) *KeyedCollection[K, T] {
	col.items.Remove(key)
	return col
}

func (col *KeyedCollection[K, T]) removeByItem(item T) *KeyedCollection[K, T] {
	col.items.RemoveByItem(item)

	return col
}

func (col *KeyedCollection[K, T]) containsKey(key K) bool {
	return col.items.ContainsKey(key)
}

func (col *KeyedCollection[K, T]) containsItem(item T) bool {
	for _, v := range col.items.values {
		if v == item {
			return true
		}
	}

	return false
}

// values returns a slice of all items in the collection.
func (col *KeyedCollection[K, T]) values() []T {
	return col.items.Values()
}

func (col *KeyedCollection[K, T]) where(predicate func(key K, item T) bool) *KeyedCollection[K, T] {
	items := col.newMap()

	for _, v := range col.items.entries {
		if predicate(v.Key, v.Value) {
			items[v.Key] = &v.Value
		}
	}

	return col.newWith(items)
}

func (col *KeyedCollection[K, T]) count(predicate func(key K, item T) bool) int {
	count := 0

	for _, v := range col.items.entries {
		if predicate(v.Key, v.Value) {
			count++
		}
	}

	return count
}

func (col *KeyedCollection[K, T]) any(predicate func(key K, item T) bool) bool {
	for _, v := range col.items.entries {
		if predicate(v.Key, v.Value) {
			return true
		}
	}

	return false
}

func (col *KeyedCollection[K, T]) all(predicate func(key K, item T) bool) bool {
	for _, v := range col.items.entries {
		if !predicate(v.Key, v.Value) {
			return false
		}
	}

	return true
}

func (col *KeyedCollection[K, T]) firstWhere(predicate func(key K, item T) bool) *T {
	for _, v := range col.items.entries {
		if predicate(v.Key, v.Value) {
			return &v.Value
		}
	}

	return nil
}

func (col *KeyedCollection[K, T]) lastWhere(predicate func(key K, item T) bool) *T {
	items := col.newMap()

	for i, _ := range col.items.entries {
		v := col.items.entries[col.items.Len()-1-i]

		if predicate(v.Key, v.Value) {
			items[v.Key] = &v.Value
		}
	}

	return col.newWith(items).Last()
}

func (col *KeyedCollection[K, T]) mapFunc(mapFunc func(key K, item T) T) *KeyedCollection[K, T] {
	items := col.newMap()

	for _, v := range col.items.entries {
		mapResult := mapFunc(v.Key, v.Value)
		items[v.Key] = &mapResult
	}

	return col.newWith(items)
}

func (col *KeyedCollection[K, T]) forEach(forEachFunc func(key K, item T)) *KeyedCollection[K, T] {
	for _, v := range col.items.entries {
		forEachFunc(v.Key, v.Value)
	}

	return col
}
