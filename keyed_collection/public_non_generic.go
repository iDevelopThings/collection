package collection

func (col *KeyedCollection[K, T]) newKeyType(key any) *K {
	newKey := (*K)(nil)
	if key != nil {
		newKey = key.(*K)
	}

	return newKey
}

func (col *KeyedCollection[K, T]) newValueType(value any) *T {
	newValue := (*T)(nil)
	if value != nil {
		newValue = value.(*T)
	}

	return newValue
}

func (col *KeyedCollection[K, T]) NGAdd(key any, item any) NonGenericKeyedCollection {
	col.Lock()
	defer col.Unlock()

	newKey := col.newKeyType(key)
	newItem := col.newValueType(item)

	col.items.Add(*newKey, *newItem)

	return col
}
func (col *KeyedCollection[K, T]) NGRemoveByKey(key any) NonGenericKeyedCollection {
	col.Lock()
	defer col.Unlock()

	return col.removeByKey(*col.newKeyType(key))
}
func (col *KeyedCollection[K, T]) NGRemoveByItem(item any) NonGenericKeyedCollection {
	col.Lock()
	defer col.Unlock()

	newItem := (*T)(nil)
	if item != nil {
		newItem = item.(*T)
	}

	return col.removeByItem(*newItem)
}
func (col *KeyedCollection[K, T]) NGContainsByKey(key any) bool {
	col.Lock()
	defer col.Unlock()

	return col.containsKey(*col.newKeyType(key))
}
func (col *KeyedCollection[K, T]) NGContainsByItem(item any) bool {
	col.Lock()
	defer col.Unlock()

	newItem := (*T)(nil)
	if item != nil {
		newItem = item.(*T)
	}

	return col.containsItem(*newItem)
}
func (col *KeyedCollection[K, T]) NGItems() []any {
	col.Lock()
	defer col.Unlock()

	items := make([]any, 0)
	for _, v := range col.items.entries {
		items = append(items, v.Value)
	}
	return items
}
func (col *KeyedCollection[K, T]) NGFirst() any {
	col.Lock()
	defer col.Unlock()

	for _, v := range col.items.entries {
		return v.Value
	}

	return nil
}

func (col *KeyedCollection[K, T]) NGLast() any {
	col.Lock()
	defer col.Unlock()

	var last T
	for _, v := range col.items.entries {
		last = v.Value
	}

	return last
}
