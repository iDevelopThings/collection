# Collection

[![Go Report Card](https://goreportcard.com/badge/github.com/idevelopthings/collection)](https://goreportcard.com/report/github.com/idevelopthings/collection)

`Collection` is a package for handling collections in Go. It provides a set of utilities for creating and manipulating collections of structs in an intuitive way.

## Installation

Use `go get` to install the package.

```sh
go get github.com/idevelopthings/collection
```

## Usage

First, import the package into your project:

```go
import "github.com/idevelopthings/collection"
```

Next, you'll need to define the type of your item:

```go
type MyItem struct {
    Name string
}
```

Create a new collection of your type:

```go
items := collection.New[*MyItem]()
// Or pass items:
items := collection.New[*MyItem](
	&MyItem{Name: "one"},
	&MyItem{Name: "two"},
)
```

Now you can manipulate the collection with the methods provided by the package:

### Add

Add an item to the collection:

```go
items.Add(&MyItem{Name:"Three"}) // returns *collections.Collection[*MyItem]
```

### Remove

Remove an item from the collection:

```go
items.Remove(item MyItem) // returns *collections.Collection[*MyItem]
```

### Contains

Check if the collection contains a certain item:

```go
items.Contains(item MyItem) // returns bool
```

### Items

Retrieve all items in the collection:

```go
items.Items() // returns []*MyItem
```

### Len

Get the number of items in the collection:

```go
items.Len() // returns int
```

### IsEmpty

Check if the collection is empty:

```go
items.IsEmpty() // returns bool
```

### IsNotEmpty

Check if the collection is not empty:

```go
items.IsNotEmpty() // returns bool
```

### First

Get the first item in the collection:

```go
items.First() // returns *MyItem
```

### Last

Get the last item in the collection:

```go
items.Last() // returns *MyItem
```

### AddRange

Add a range of items to the collection:

```go
items.AddRange(items []MyItem) // returns *collections.Collection[*MyItem]
```

### Clear

Remove all items from the collection:

```go
items.Clear() // returns *collections.Collection[*MyItem]
```

### Where

Find all items that satisfy a certain predicate:

```go
items.Where(predicate func(item MyItem) bool) // returns *collections.Collection[*MyItem]
```

### Count

Count the number of items that satisfy a certain predicate:

```go
items.Count(predicate func(item MyItem) bool) // returns int
```

### Any

Check if any item in the collection satisfies a certain predicate:

```go
items.Any(predicate func(item MyItem) bool) // returns bool
```

### All

Check if all items in the collection satisfy a certain predicate:

```go
items.All(predicate func(item MyItem) bool) // returns bool
```

### FirstWhere

Find the first item that satisfies a certain predicate:

```go
items.FirstWhere(predicate func(item MyItem) bool) // returns **MyItem
```

### LastWhere

Find the last item that satisfies a certain predicate:

```go
items.LastWhere(predicate func(item MyItem) bool) // returns **MyItem
```

### Map

Transform the collection using a mapping function & return a new collection(so the original isn't modified)

```go
items.Map(mapFunc func(index int, item MyItem) MyItem) // returns *collections.Collection[*MyItem]
```

### ForEach

Iterate over the collection:

```go
items.ForEach(forEachFunc func(index int, item MyItem)) // returns *collections.Collection[*MyItem]
```

### Set

Set the value of an item at a specific index:

```go
items.Set(index int, value MyItem) // returns *collections.Collection[*MyItem]
```

### IndexOf

Find the index of an item:

```go
items.IndexOf(item MyItem) // returns int
```

### Filter

Filter the collection based on a predicate function:

```go
items.Filter(filterFunc func(index int, item MyItem) bool) // returns *collections.Collection[*MyItem]
```

## License

MIT
