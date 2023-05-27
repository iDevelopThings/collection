package collection

func (col *Collection[T]) get(index int) *T {
	return col.items[index]
}

func (col *Collection[T]) add(item T) *Collection[T] {
	col.items = append(col.items, &item)

	return col
}

func (col *Collection[T]) remove(item T) *Collection[T] {
	for i, v := range col.items {
		if *v == item {
			col.items = append(col.items[:i], col.items[i+1:]...)
			return col
		}
	}

	return col
}

func (col *Collection[T]) contains(item T) bool {
	for _, v := range col.items {
		if *v == item {
			return true
		}
	}

	return false
}

func (col *Collection[T]) where(predicate func(item T) bool) []*T {
	items := make([]*T, 0)

	for _, v := range col.items {
		if predicate(*v) {
			items = append(items, v)
		}
	}

	return items
}

func (col *Collection[T]) count(predicate func(item T) bool) int {
	count := 0

	for _, v := range col.items {
		if predicate(*v) {
			count++
		}
	}

	return count
}

func (col *Collection[T]) any(predicate func(item T) bool) bool {
	for _, v := range col.items {
		if predicate(*v) {
			return true
		}
	}

	return false
}

func (col *Collection[T]) all(predicate func(item T) bool) bool {
	for _, v := range col.items {
		if !predicate(*v) {
			return false
		}
	}

	return true
}

func (col *Collection[T]) firstWhere(predicate func(item T) bool) *T {
	for _, v := range col.items {
		if predicate(*v) {
			return v
		}
	}

	return nil
}

func (col *Collection[T]) lastWhere(predicate func(item T) bool) *T {
	for i := len(col.items) - 1; i >= 0; i-- {
		if predicate(*col.items[i]) {
			return col.items[i]
		}
	}

	return nil
}

func (col *Collection[T]) mapFunc(mapFunc func(index int, item T) T) *Collection[T] {
	for i, v := range col.items {
		mapResult := mapFunc(i, *v)
		col.items[i] = &mapResult
	}

	return col
}

func (col *Collection[T]) forEach(forEachFunc func(index int, item T)) *Collection[T] {
	for i, v := range col.items {
		forEachFunc(i, *v)
	}

	return col
}

/*// keyByFunc The keyBy method keys the collection by the given key.
// If multiple items have the same key, only the last one will appear in the new collection:
func (col *Collection[T]) keyByFunc(mapFunc func(index int, item T) any) any {
	keys := make([]any, 0)
	// A map of collection index to key result
	entries := make(map[int]any, 0)

	var firstKeyType reflect.Type

	for i, item := range col.items {
		keyResult := mapFunc(i, *item)
		if !arrContains(keys, keyResult) {
			keys = append(keys, keyResult)
			if firstKeyType == nil {
				firstKeyType = reflect.TypeOf(keyResult)
			}
		}

		if reflect.TypeOf(keyResult) != firstKeyType {
			panic("keyByFunc: All keys must be of the same type")
		}

		entries[i] = keyResult
	}

	// Now we have a list of keys and a map of entries
	// We'll create a new keyed collection and add the entries to it

	keyedCollection := collection.NewFromTypes(firstKeyType, col.itemType)

	for i, key := range entries {
		keyedCollection.NGAdd(key, *col.items[i])
	}

	return keyedCollection
}

func (col *Collection[T]) keyByProperty(prop string) any {
	keys := make([]any, 0)
	// A map of collection index to key result
	entries := make(map[int]any, 0)

	var firstKeyType reflect.Type

	if col.itemType.Kind() != reflect.Struct {
		panic("keyByProperty: Collection must be of struct type")
	}

	for i, item := range col.items {
		propValue := reflect.ValueOf(item).Elem().FieldByName(prop).Interface()
		if !arrContains(keys, propValue) {
			keys = append(keys, propValue)
			if firstKeyType == nil {
				firstKeyType = reflect.TypeOf(propValue)
			}
		}

		if reflect.TypeOf(propValue) != firstKeyType {
			panic("keyByProperty: All keys must be of the same type")
		}

		entries[i] = propValue
	}

	// Now we have a list of keys and a map of entries
	// We'll create a new keyed collection and add the entries to it

	keyedCollection := collection.NewFromTypes(firstKeyType, col.itemType)

	for i, key := range entries {
		keyedCollection.NGAdd(key, *col.items[i])
	}

	return keyedCollection
}*/
