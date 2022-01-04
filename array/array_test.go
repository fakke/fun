package array

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestOfLen(t *testing.T) {
	assert.Equal(t, OfLen[int](0).Len(), 0)
	assert.Equal(t, OfLen[int](1).Len(), 1)
	assert.Equal(t, OfLen[int](2).Len(), 2)
	assert.Equal(t, OfLen[int](3).Len(), 3)
	assert.Equal(t, OfLen[int](4).Len(), 4)
	assert.Equal(t, OfLen[int](0).Cap(), 0)
	assert.Equal(t, OfLen[int](1).Cap(), 1)
	assert.Equal(t, OfLen[int](2).Cap(), 2)
	assert.Equal(t, OfLen[int](3).Cap(), 3)
	assert.Equal(t, OfLen[int](4).Cap(), 4)
}

func TestOfCap(t *testing.T) {
	assert.Equal(t, OfCap[int](0).Len(), 0)
	assert.Equal(t, OfCap[int](1).Len(), 0)
	assert.Equal(t, OfCap[int](2).Len(), 0)
	assert.Equal(t, OfCap[int](3).Len(), 0)
	assert.Equal(t, OfCap[int](4).Len(), 0)
	assert.Equal(t, OfCap[int](0).Cap(), 0)
	assert.Equal(t, OfCap[int](1).Cap(), 1)
	assert.Equal(t, OfCap[int](2).Cap(), 2)
	assert.Equal(t, OfCap[int](3).Cap(), 3)
	assert.Equal(t, OfCap[int](4).Cap(), 4)
}

func TestOfSize(t *testing.T) {
	assert.Equal(t, OfSize[int](0, 0).Len(), 0)
	assert.Equal(t, OfSize[int](0, 1).Len(), 0)
	assert.Equal(t, OfSize[int](0, 2).Len(), 0)
	assert.Equal(t, OfSize[int](0, 3).Len(), 0)
	assert.Equal(t, OfSize[int](0, 4).Len(), 0)
	assert.Equal(t, OfSize[int](0, 0).Cap(), 0)
	assert.Equal(t, OfSize[int](0, 1).Cap(), 1)
	assert.Equal(t, OfSize[int](0, 2).Cap(), 2)
	assert.Equal(t, OfSize[int](0, 3).Cap(), 3)
	assert.Equal(t, OfSize[int](0, 4).Cap(), 4)
}

func TestOf(t *testing.T) {
	assert.Equal(t, Of(10, 11, 22).Len(), 3)
	assert.Equal(t, Of(10, 11, 22).At(0), 10)
	assert.Equal(t, Of(10, 11, 22).At(1), 11)
	assert.Equal(t, Of(10, 11, 22).At(2), 22)
}

func TestArray_Append(t *testing.T) {
	a := Of(0, 1)
	a.Append(2)
	b := Of(0, 1, 2)
	assert.Equal(t, a, b)
}

func TestArray_AppendAll(t *testing.T) {
	a := Of(0, 1)
	a.AppendAll(2, 3)
	b := Of(0, 1, 2, 3)
	assert.Equal(t, a, b)
}

func TestArray_AppendFrom(t *testing.T) {
	a := Of(0, 1)
	b := Of(2, 3)
	a.AppendFrom(b)
	c := Of(0, 1, 2, 3)
	assert.Equal(t, a, c)
}

func TestArray_InsertAt(t *testing.T) {
	a := Of(0, 1, 3, 4)
	a.InsertAt(2)(2)
	assert.Equal(t, a.At(2), 2)
}

func TestAppend(t *testing.T) {
	a := Of[int](0, 1, 2)
	b := Of[int](3, 4, 5)
	a.AppendFrom(b)
	assert.Equal(t, a, Of(0, 1, 2, 3, 4, 5))
}

func TestMap(t *testing.T) {
	ints := Of[int](0, 1, 2)
	strings := Map[int, string](ints, strconv.Itoa)
	assert.Equal(t, strings, Of("0", "1", "2"))
}
