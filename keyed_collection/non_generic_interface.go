package collection

type NonGenericKeyedCollection interface {
	// Add adds an item to the collection.
	NGAdd(key any, item any) NonGenericKeyedCollection
	// Remove removes an item from the collection by key.
	NGRemoveByKey(key any) NonGenericKeyedCollection
	// Remove removes an item from the collection by item.
	NGRemoveByItem(item any) NonGenericKeyedCollection
	// ContainsByKey returns true if the collection contains the key.
	NGContainsByKey(key any) bool
	// ContainsByItem returns true if the collection contains the item.
	NGContainsByItem(item any) bool
	// Items returns the items in the collection.
	NGItems() []any
	// First returns the first item in the collection.
	NGFirst() any
	// Last returns the last item in the collection.
	NGLast() any
}
