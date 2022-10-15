package array

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestOfLen(t *testing.T) {
	assert.Equal(t, 0, OfLen[int](0).Len())
	assert.Equal(t, 1, OfLen[int](1).Len())
	assert.Equal(t, 2, OfLen[int](2).Len())
	assert.Equal(t, 3, OfLen[int](3).Len())
	assert.Equal(t, 4, OfLen[int](4).Len())
	assert.Equal(t, 0, OfLen[int](0).Cap())
	assert.Equal(t, 1, OfLen[int](1).Cap())
	assert.Equal(t, 2, OfLen[int](2).Cap())
	assert.Equal(t, 3, OfLen[int](3).Cap())
	assert.Equal(t, 4, OfLen[int](4).Cap())
}

func TestOfCap(t *testing.T) {
	assert.Equal(t, 0, OfCap[int](0).Len())
	assert.Equal(t, 0, OfCap[int](1).Len())
	assert.Equal(t, 0, OfCap[int](2).Len())
	assert.Equal(t, 0, OfCap[int](3).Len())
	assert.Equal(t, 0, OfCap[int](4).Len())
	assert.Equal(t, 0, OfCap[int](0).Cap())
	assert.Equal(t, 1, OfCap[int](1).Cap())
	assert.Equal(t, 2, OfCap[int](2).Cap())
	assert.Equal(t, 3, OfCap[int](3).Cap())
	assert.Equal(t, 4, OfCap[int](4).Cap())
}

func TestOfSize(t *testing.T) {
	assert.Equal(t, 0, OfSize[int](0, 0).Len())
	assert.Equal(t, 0, OfSize[int](0, 1).Len())
	assert.Equal(t, 0, OfSize[int](0, 2).Len())
	assert.Equal(t, 0, OfSize[int](0, 3).Len())
	assert.Equal(t, 0, OfSize[int](0, 4).Len())
	assert.Equal(t, 0, OfSize[int](0, 0).Cap())
	assert.Equal(t, 1, OfSize[int](0, 1).Cap())
	assert.Equal(t, 2, OfSize[int](0, 2).Cap())
	assert.Equal(t, 3, OfSize[int](0, 3).Cap())
	assert.Equal(t, 4, OfSize[int](0, 4).Cap())
}

func TestOf(t *testing.T) {
	assert.Equal(t, 3, Of(10, 11, 22).Len())
	assert.Equal(t, 10, Of(10, 11, 22).At(0))
	assert.Equal(t, 11, Of(10, 11, 22).At(1))
	assert.Equal(t, 22, Of(10, 11, 22).At(2))
}

func TestArray_Append(t *testing.T) {
	a := Of(0, 1)
	a.Append(2)
	b := Of(0, 1, 2)
	assert.Equal(t, b, a)
}

func TestArray_AppendAll(t *testing.T) {
	a := Of(0, 1)
	a.AppendAll(2, 3)
	b := Of(0, 1, 2, 3)
	assert.Equal(t, b, a)
}

func TestArray_AppendFrom(t *testing.T) {
	a := Of(0, 1)
	b := Of(2, 3)
	a.AppendFrom(b)
	assert.Equal(t, Of(0, 1, 2, 3), a)
}

func TestArray_InsertAt(t *testing.T) {
	a := Of(0, 1, 3, 4)
	a.InsertAt(2)(2)
	assert.Equal(t, 2, a.At(2))
}

func TestAppend(t *testing.T) {
	a := Of(0, 1, 2)
	b := Of(3, 4, 5)
	a.AppendFrom(b)
	assert.Equal(t, Of(0, 1, 2, 3, 4, 5), a)
}

func TestMap(t *testing.T) {
	ints := Of(0, 1, 2)
	strings := Map(strconv.Itoa)(ints)
	assert.Equal(t, Of("0", "1", "2"), strings)
}
