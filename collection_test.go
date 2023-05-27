package collection

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CollectionTestSuite struct {
	suite.Suite
}

func (suite *CollectionTestSuite) TestCreate() {
	col := New[int]()

	suite.NotNil(col)
	suite.NotNil(col.items)
	suite.Equal(0, len(col.items))
}

func (suite *CollectionTestSuite) TestCreateWithItems() {
	col := New[int](1, 2, 3)

	suite.NotNil(col)
	suite.NotNil(col.items)
	suite.Equal(3, len(col.items))
}

func (suite *CollectionTestSuite) TestAdd() {
	col := New[int]()

	col.Add(1)
	col.Add(2)
	col.Add(3)

	suite.Equal(3, len(col.items))
}

func (suite *CollectionTestSuite) TestRemove() {
	col := New[int](1, 2, 3)

	col.Remove(2)

	suite.Equal(2, len(col.items))
}

func (suite *CollectionTestSuite) TestContains() {
	col := New[int](1, 2, 3)

	suite.True(col.Contains(2))
}

func (suite *CollectionTestSuite) TestItems() {
	col := New[int](1, 2, 3)

	items := col.Items()

	suite.Equal(3, len(items))
}

func (suite *CollectionTestSuite) TestLen() {
	col := New[int](1, 2, 3)

	suite.Equal(3, col.Len())
}

func (suite *CollectionTestSuite) TestIsEmpty() {
	col := New[int]()
	suite.True(col.IsEmpty())
}

func (suite *CollectionTestSuite) TestIsNotEmpty() {
	col := New[int](1, 2, 3)
	suite.True(col.IsNotEmpty())
}

func (suite *CollectionTestSuite) TestFirst() {
	col := New[int](1, 2, 3)
	suite.Equal(1, *col.First())
}

func (suite *CollectionTestSuite) TestLast() {
	col := New[int](1, 2, 3)
	suite.Equal(3, *col.Last())
}

func (suite *CollectionTestSuite) TestWhere() {
	col := New[int](1, 2, 3, 4, 5)
	suite.Equal(2, len(col.Where(func(item int) bool { return item%2 == 0 })))
}

func (suite *CollectionTestSuite) TestAny() {
	col := New[int](1, 2, 3, 4, 5)
	suite.True(col.Any(func(item int) bool { return item%2 == 0 }))
}

func (suite *CollectionTestSuite) TestAll() {
	col := New[int](1, 2, 3, 4, 5)
	suite.False(col.All(func(item int) bool { return item%2 == 0 }))
}

func (suite *CollectionTestSuite) TestMap() {
	col := New[int](1, 2, 3, 4, 5)
	mappedCol := col.Map(func(index int, item int) int { return item * 2 })
	suite.Equal(2, *mappedCol.First())
	suite.Equal(10, *mappedCol.Last())
}

func (suite *CollectionTestSuite) TestFirstWhere() {
	col := New[int](1, 2, 3, 4, 5)
	suite.Equal(2, *col.FirstWhere(func(item int) bool { return item%2 == 0 }))
}

func (suite *CollectionTestSuite) TestLastWhere() {
	col := New[int](1, 2, 3, 4, 5)
	suite.Equal(4, *col.LastWhere(func(item int) bool { return item%2 == 0 }))
}

func (suite *CollectionTestSuite) TestForEach() {
	col := New[int](1, 2, 3, 4, 5)
	col.ForEach(func(index int, item int) {
		i := item * 2
		col.items[index] = &i
	})

	suite.Equal(2, *col.First())
	suite.Equal(10, *col.Last())
}

func (suite *CollectionTestSuite) TestKeyByFunc() {
	type MyTestingItem struct {
		ID           int
		TestProperty string
	}
	col := New[MyTestingItem](
		MyTestingItem{1, "one"},
		MyTestingItem{2, "two"},
		MyTestingItem{3, "three"},
	)

	testPropCollection := KeyBy[string, MyTestingItem](col, func(index int, item MyTestingItem) string {
		return item.TestProperty
	})

	testIdCollection := KeyBy[int, MyTestingItem](col, func(index int, item MyTestingItem) int {
		return item.ID
	})

	suite.True(testPropCollection.ContainsKey("one"))
	suite.True(testPropCollection.ContainsKey("two"))
	suite.True(testPropCollection.ContainsKey("three"))

	suite.Equal(1, testPropCollection.GetByKey("one").ID)
	suite.Equal(2, testPropCollection.GetByKey("two").ID)
	suite.Equal(3, testPropCollection.GetByKey("three").ID)

	suite.True(testIdCollection.ContainsKey(1))
	suite.True(testIdCollection.ContainsKey(2))
	suite.True(testIdCollection.ContainsKey(3))

	suite.Equal("one", testIdCollection.GetByKey(1).TestProperty)
	suite.Equal("two", testIdCollection.GetByKey(2).TestProperty)
	suite.Equal("three", testIdCollection.GetByKey(3).TestProperty)
}

func TestCollectionTestSuite(t *testing.T) {
	suite.Run(t, new(CollectionTestSuite))
}
