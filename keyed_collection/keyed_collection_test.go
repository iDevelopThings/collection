package collection

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"
)

type KeyedCollectionTestSuite struct {
	suite.Suite
}

func (suite *KeyedCollectionTestSuite) TestCreate() {
	col := New[string, int]()

	suite.NotNil(col)
	suite.NotNil(col.items)
	suite.Equal(0, len(col.items.entries))
}

func (suite *KeyedCollectionTestSuite) TestCreateWithItems() {
	col := New[string, int](map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})

	suite.NotNil(col)
	suite.NotNil(col.items)
	suite.Equal(3, len(col.items.entries))
}

func (suite *KeyedCollectionTestSuite) TestAdd() {
	col := New[string, int]()

	col.Add("one", 1)
	col.Add("two", 2)
	col.Add("three", 3)

	suite.Equal(3, len(col.items.entries))
}

func (suite *KeyedCollectionTestSuite) TestRemove() {
	col := New[string, int](map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})

	col.RemoveByKey("two")
	col.RemoveByItem(3)

	suite.Equal(1, len(col.items.entries))
	suite.Equal(1, *col.items.Get("one"))
}

func (suite *KeyedCollectionTestSuite) TestContains() {
	col := New[string, int](map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})

	suite.True(col.ContainsKey("two"))
	suite.True(col.ContainsItem(3))
}

func (suite *KeyedCollectionTestSuite) TestItems() {
	col := New[string, int](map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})

	items := col.Items()

	suite.Equal(3, len(items))
	suite.True(*items["one"] == 1)
}

func (suite *KeyedCollectionTestSuite) TestLen() {
	col := New[string, int](map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})

	suite.Equal(3, col.Len())
}

func (suite *KeyedCollectionTestSuite) TestIsEmpty() {
	col := New[string, int]()
	suite.True(col.IsEmpty())
}

func (suite *KeyedCollectionTestSuite) TestIsNotEmpty() {
	col := New[string, int](map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})
	suite.True(col.IsNotEmpty())
}

func (suite *KeyedCollectionTestSuite) TestFirst() {
	col := New[string, int](map[string]int{
		"one": 1,
	})

	suite.Equal(1, *col.First())
}

func (suite *KeyedCollectionTestSuite) TestLast() {
	col := New[string, int](map[string]int{
		"three": 3,
	})
	suite.Equal(3, *col.Last())
}

func (suite *KeyedCollectionTestSuite) TestWhere() {
	col := New[string, int](map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	})

	whereCol := col.Where(func(key string, item int) bool { return item%2 == 0 })

	suite.Equal(2, whereCol.Len())
}

func (suite *KeyedCollectionTestSuite) TestAny() {
	col := New[string, int](map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	})
	suite.True(col.Any(func(key string, item int) bool { return item%2 == 0 }))
}

func (suite *KeyedCollectionTestSuite) TestAll() {
	col := New[string, int](map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	})
	suite.False(col.All(func(key string, item int) bool { return item%2 == 0 }))
}

/*func (suite *KeyedCollectionTestSuite) TestMap() {
	col := New[string, int](map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	})
	mappedCollection := col.Map(func(key string, item int) int { return item * 2 })
	suite.Equal(2, *mappedCollection.First())
	suite.Equal(10, *mappedCollection.Last())
}*/

func (suite *KeyedCollectionTestSuite) TestFirstWhere() {
	col := New[string, int](map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	})
	suite.Equal(2, *col.FirstWhere(func(key string, item int) bool { return item%2 == 0 }))
}

func (suite *KeyedCollectionTestSuite) TestLastWhere() {
	col := New[string, int](map[string]int{
		"one":       1,
		"one_again": 1,
		"three":     3,
		"four":      4,
		"five":      5,
	})
	last := col.LastWhere(func(key string, item int) bool {
		return item == 1
	})

	suite.Equal(1, *last)
}

func (suite *KeyedCollectionTestSuite) TestForEach() {
	col := New[string, int](map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	})

	items := make([]int, 0)
	col.ForEach(func(key string, item int) {
		items = append(items, item*2)
	})

	sort.Ints(items)

	for i, item := range items {
		suite.Equal((i+1)*2, item)
	}
}

func TestKeyedCollectionTestSuite(t *testing.T) {
	suite.Run(t, new(KeyedCollectionTestSuite))
}
