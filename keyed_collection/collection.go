package collection

import (
	"reflect"
	"sync"

	"golang.org/x/exp/constraints"
)

type KeyedCollection[K constraints.Ordered, T comparable] struct {
	*sync.Mutex

	items *SortedMap[K, T]

	keyType  reflect.Type
	itemType reflect.Type
}

func New[K constraints.Ordered, T comparable](items ...map[K]T) *KeyedCollection[K, T] {
	col := &KeyedCollection[K, T]{
		Mutex: &sync.Mutex{},
	}

	col.items = NewSortedMapOf[K, T]()
	col.keyType = reflect.TypeOf((*K)(nil)).Elem()
	col.itemType = reflect.TypeOf((*T)(nil)).Elem()

	if len(items) > 0 {
		for k, t := range items[0] {
			col.items.Add(k, t)
		}
	}

	return col
}

func (col *KeyedCollection[K, T]) newMap() map[K]*T {
	return make(map[K]*T, 0)
}
func (col *KeyedCollection[K, T]) newWith(items map[K]*T) *KeyedCollection[K, T] {
	newCol := New[K, T]()
	newCol.items = NewSortedMapOf[K, T]()
	if len(items) > 0 {
		for k, t := range items {
			newCol.items.Add(k, *t)
		}
	}
	return newCol
}

/*func (col *KeyedCollection[K, T]) updateKeys(method string, values ...map[string]any) {
	if method == "add" {
		keyAdded := values[0]["key"].(K)

		for _, key := range col.keys {
			if key == keyAdded {
				return
			}
		}

		col.keys = append(col.keys, keyAdded)
		return
	}

	if method == "remove" {
		keyRemoved := values[0]["key"].(K)

		col.keysLock.Lock()
		defer col.keysLock.Unlock()

		for i, key := range col.keys {
			if key == keyRemoved {
				col.keys = append(col.keys[:i], col.keys[i+1:]...)
				return
			}
		}

		return
	}

}*/

// func NewFromTypes(keyType, valueType reflect.Type) NonGenericKeyedCollection {
// 	val := reflect.New(valueType).Elem().Interface()
//
// 	// Get the type of the keyed collection
// 	keyedCollectionType := reflect.TypeOf((*KeyedCollection[string, val])(nil)).Elem()
// 	keyedCollectionType = keyedCollectionType.Elem()
// 	// Assign the key and value types to the keyed collection type
//
// 	// Create a new instance of the keyed collection type, with the given key and value types
// 	keyedCollection := reflect.New(keyedCollectionType).Elem()
// 	itemsMap := reflect.MakeMap(reflect.MapOf(keyType, valueType))
// 	keyedCollection.Set(itemsMap)
//
// 	return keyedCollection.Interface().(NonGenericKeyedCollection)
// }
