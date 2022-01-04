package slice

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var isEven = func(x int) bool {
	return x%2 == 0
}
var isOdd = func(x int) bool {
	return !isEven(x)
}
var iadd = func(a, b int) int {
	return a + b
}
var imul = func(a, b int) int {
	return a + b
}
var appendInt = func(xs []int, x int) []int {
	return append(xs, x)
}
var intAppend = func(x int, xs []int) []int {
	return append(xs, x)
}
var itoa = func(x int) string {
	return strconv.Itoa(x)
}
var strlen = func(s string) int {
	return len(s)
}
var istrlen = func(x int) int {
	return len(itoa(x))
}

func TestIndex(t *testing.T) {
	assert.Equal(t, Index([]int{}, isOdd), -1)
	assert.Equal(t, Index([]int{0}, isOdd), -1)
	assert.Equal(t, Index([]int{1}, isOdd), 0)
	assert.Equal(t, Index([]int{42}, isOdd), -1)
	assert.Equal(t, Index([]int{2, 4, 6}, isOdd), -1)
	assert.Equal(t, Index([]int{1, 4, 6}, isOdd), 0)
	assert.Equal(t, Index([]int{2, 1, 6}, isOdd), 1)
	assert.Equal(t, Index([]int{2, 4, 1}, isOdd), 2)
}

func TestIndexOf(t *testing.T) {
	assert.Equal(t, IndexOf([]int{}, 0), -1)
	assert.Equal(t, IndexOf([]int{0}, 0), 0)
	assert.Equal(t, IndexOf([]int{1, 2, 3}, 0), -1)
	assert.Equal(t, IndexOf([]int{1, 2, 3}, 1), 0)
	assert.Equal(t, IndexOf([]int{1, 2, 3}, 2), 1)
	assert.Equal(t, IndexOf([]int{1, 2, 3}, 3), 2)
}

func TestFoldL(t *testing.T) {
	// empty
	assert.Equal(t, FoldL([]int{}, 0, iadd), 0)
	assert.Equal(t, FoldL([]int{}, 1, iadd), 1)
	assert.Equal(t, FoldL([]int{}, 42, iadd), 42)
	assert.Equal(t, FoldL([]int{}, -1, iadd), -1)
	assert.Equal(t, FoldL([]int{}, -42, iadd), -42)
	// zero
	assert.Equal(t, FoldL([]int{0}, 0, iadd), 0)
	assert.Equal(t, FoldL([]int{0}, 1, iadd), 1)
	assert.Equal(t, FoldL([]int{0}, 42, iadd), 42)
	assert.Equal(t, FoldL([]int{0}, -1, iadd), -1)
	assert.Equal(t, FoldL([]int{0}, -42, iadd), -42)
	// several zeroes
	assert.Equal(t, FoldL([]int{0, 0}, 0, iadd), 0)
	assert.Equal(t, FoldL([]int{0, 0}, 1, iadd), 1)
	assert.Equal(t, FoldL([]int{0, 0}, 42, iadd), 42)
	assert.Equal(t, FoldL([]int{0, 0}, -1, iadd), -1)
	assert.Equal(t, FoldL([]int{0, 0}, -42, iadd), -42)
	// one value
	assert.Equal(t, FoldL([]int{42}, 0, iadd), 42)
	assert.Equal(t, FoldL([]int{42}, 1, iadd), 43)
	assert.Equal(t, FoldL([]int{42}, 42, iadd), 84)
	assert.Equal(t, FoldL([]int{42}, -1, iadd), 41)
	assert.Equal(t, FoldL([]int{42}, -42, iadd), 0)
	// several values
	assert.Equal(t, FoldL([]int{1, 2, 3}, 0, iadd), 6)
	assert.Equal(t, FoldL([]int{1, 2, 3}, 1, iadd), 7)
	assert.Equal(t, FoldL([]int{1, 2, 3}, 42, iadd), 48)
	assert.Equal(t, FoldL([]int{1, 2, 3}, -1, iadd), 5)
	assert.Equal(t, FoldL([]int{1, 2, 3}, -42, iadd), -36)
	// non-commutative
	assert.Equal(t, FoldL([]int{}, []int{}, appendInt), []int{})
	assert.Equal(t, FoldL([]int{}, []int{0}, appendInt), []int{0})
	assert.Equal(t, FoldL([]int{0}, []int{}, appendInt), []int{0})
	assert.Equal(t, FoldL([]int{2}, []int{1}, appendInt), []int{1, 2})
	assert.Equal(t, FoldL([]int{3, 4}, []int{1, 2}, appendInt), []int{1, 2, 3, 4})
}

func TestFoldR(t *testing.T) {
	// empty
	assert.Equal(t, FoldR([]int{}, 0, iadd), 0)
	assert.Equal(t, FoldR([]int{}, 1, iadd), 1)
	assert.Equal(t, FoldR([]int{}, 42, iadd), 42)
	assert.Equal(t, FoldR([]int{}, -1, iadd), -1)
	assert.Equal(t, FoldR([]int{}, -42, iadd), -42)
	// zero
	assert.Equal(t, FoldR([]int{0}, 0, iadd), 0)
	assert.Equal(t, FoldR([]int{0}, 1, iadd), 1)
	assert.Equal(t, FoldR([]int{0}, 42, iadd), 42)
	assert.Equal(t, FoldR([]int{0}, -1, iadd), -1)
	assert.Equal(t, FoldR([]int{0}, -42, iadd), -42)
	// several zeroes
	assert.Equal(t, FoldR([]int{0, 0}, 0, iadd), 0)
	assert.Equal(t, FoldR([]int{0, 0}, 1, iadd), 1)
	assert.Equal(t, FoldR([]int{0, 0}, 42, iadd), 42)
	assert.Equal(t, FoldR([]int{0, 0}, -1, iadd), -1)
	assert.Equal(t, FoldR([]int{0, 0}, -42, iadd), -42)
	// one value
	assert.Equal(t, FoldR([]int{42}, 0, iadd), 42)
	assert.Equal(t, FoldR([]int{42}, 1, iadd), 43)
	assert.Equal(t, FoldR([]int{42}, 42, iadd), 84)
	assert.Equal(t, FoldR([]int{42}, -1, iadd), 41)
	assert.Equal(t, FoldR([]int{42}, -42, iadd), 0)
	// several values
	assert.Equal(t, FoldR([]int{1, 2, 3}, 0, iadd), 6)
	assert.Equal(t, FoldR([]int{1, 2, 3}, 1, iadd), 7)
	assert.Equal(t, FoldR([]int{1, 2, 3}, 42, iadd), 48)
	assert.Equal(t, FoldR([]int{1, 2, 3}, -1, iadd), 5)
	assert.Equal(t, FoldR([]int{1, 2, 3}, -42, iadd), -36)
	// non-commutative
	assert.Equal(t, FoldR([]int{}, []int{}, intAppend), []int{})
	assert.Equal(t, FoldR([]int{}, []int{0}, intAppend), []int{0})
	assert.Equal(t, FoldR([]int{0}, []int{}, intAppend), []int{0})
	assert.Equal(t, FoldR([]int{2}, []int{1}, intAppend), []int{1, 2})
	assert.Equal(t, FoldR([]int{3, 4}, []int{1, 2}, intAppend), []int{1, 2, 4, 3})
}

func TestScanL(t *testing.T) {
	// empty
	assert.Equal(t, ScanL([]int{}, 0, iadd), []int{0})
	assert.Equal(t, ScanL([]int{}, 1, iadd), []int{1})
	assert.Equal(t, ScanL([]int{}, 42, iadd), []int{42})
	assert.Equal(t, ScanL([]int{}, -1, iadd), []int{-1})
	assert.Equal(t, ScanL([]int{}, -42, iadd), []int{-42})
	// zero
	assert.Equal(t, ScanL([]int{0}, 0, iadd), []int{0, 0})
	assert.Equal(t, ScanL([]int{0}, 1, iadd), []int{1, 1})
	assert.Equal(t, ScanL([]int{0}, 42, iadd), []int{42, 42})
	assert.Equal(t, ScanL([]int{0}, -1, iadd), []int{-1, -1})
	assert.Equal(t, ScanL([]int{0}, -42, iadd), []int{-42, -42})
	// several zeroes
	assert.Equal(t, ScanL([]int{0, 0}, 0, iadd), []int{0, 0, 0})
	assert.Equal(t, ScanL([]int{0, 0}, 1, iadd), []int{1, 1, 1})
	assert.Equal(t, ScanL([]int{0, 0}, 42, iadd), []int{42, 42, 42})
	assert.Equal(t, ScanL([]int{0, 0}, -1, iadd), []int{-1, -1, -1})
	assert.Equal(t, ScanL([]int{0, 0}, -42, iadd), []int{-42, -42, -42})
	// one value
	assert.Equal(t, ScanL([]int{42}, 0, iadd), []int{0, 42})
	assert.Equal(t, ScanL([]int{42}, 1, iadd), []int{1, 43})
	assert.Equal(t, ScanL([]int{42}, 42, iadd), []int{42, 84})
	assert.Equal(t, ScanL([]int{42}, -1, iadd), []int{-1, 41})
	assert.Equal(t, ScanL([]int{42}, -42, iadd), []int{-42, 0})
	// several values
	assert.Equal(t, ScanL([]int{1, 2, 3}, 0, iadd), []int{0, 1, 3, 6})
	assert.Equal(t, ScanL([]int{1, 2, 3}, 1, iadd), []int{1, 2, 4, 7})
	assert.Equal(t, ScanL([]int{1, 2, 3}, 42, iadd), []int{42, 43, 45, 48})
	assert.Equal(t, ScanL([]int{1, 2, 3}, -1, iadd), []int{-1, 0, 2, 5})
	assert.Equal(t, ScanL([]int{1, 2, 3}, -42, iadd), []int{-42, -41, -39, -36})
}

func TestScanR(t *testing.T) {
	// empty
	assert.Equal(t, ScanR([]int{}, 0, iadd), []int{0})
	assert.Equal(t, ScanR([]int{}, 1, iadd), []int{1})
	assert.Equal(t, ScanR([]int{}, 42, iadd), []int{42})
	assert.Equal(t, ScanR([]int{}, -1, iadd), []int{-1})
	assert.Equal(t, ScanR([]int{}, -42, iadd), []int{-42})
	// zero
	assert.Equal(t, ScanR([]int{0}, 0, iadd), []int{0, 0})
	assert.Equal(t, ScanR([]int{0}, 1, iadd), []int{1, 1})
	assert.Equal(t, ScanR([]int{0}, 42, iadd), []int{42, 42})
	assert.Equal(t, ScanR([]int{0}, -1, iadd), []int{-1, -1})
	assert.Equal(t, ScanR([]int{0}, -42, iadd), []int{-42, -42})
	// several zeroes
	assert.Equal(t, ScanR([]int{0, 0}, 0, iadd), []int{0, 0, 0})
	assert.Equal(t, ScanR([]int{0, 0}, 1, iadd), []int{1, 1, 1})
	assert.Equal(t, ScanR([]int{0, 0}, 42, iadd), []int{42, 42, 42})
	assert.Equal(t, ScanR([]int{0, 0}, -1, iadd), []int{-1, -1, -1})
	assert.Equal(t, ScanR([]int{0, 0}, -42, iadd), []int{-42, -42, -42})
	// one value
	assert.Equal(t, ScanR([]int{42}, 0, iadd), []int{42, 0})
	assert.Equal(t, ScanR([]int{42}, 1, iadd), []int{43, 1})
	assert.Equal(t, ScanR([]int{42}, 42, iadd), []int{84, 42})
	assert.Equal(t, ScanR([]int{42}, -1, iadd), []int{41, -1})
	assert.Equal(t, ScanR([]int{42}, -42, iadd), []int{0, -42})
	// several values
	assert.Equal(t, ScanR([]int{1, 2, 3}, 0, iadd), []int{6, 5, 3, 0})
	assert.Equal(t, ScanR([]int{1, 2, 3}, 1, iadd), []int{7, 6, 4, 1})
	assert.Equal(t, ScanR([]int{1, 2, 3}, 42, iadd), []int{48, 47, 45, 42})
	assert.Equal(t, ScanR([]int{1, 2, 3}, -1, iadd), []int{5, 4, 2, -1})
	assert.Equal(t, ScanR([]int{1, 2, 3}, -42, iadd), []int{-36, -37, -39, -42})
}

func TestMap(t *testing.T) {
	assert.Equal(t, Map([]int{}, itoa), []string{})
	assert.Equal(t, Map([]int{0}, itoa), []string{"0"})
	assert.Equal(t, Map([]int{42}, itoa), []string{"42"})
	assert.Equal(t, Map([]int{1, 2, 3}, itoa), []string{"1", "2", "3"})
	assert.Equal(t, Map([]int{1, 2, 3}, istrlen), []int{1, 1, 1})
	assert.Equal(t, Map([]int{1, 2, 3}, istrlen), []int{1, 1, 1})
	assert.Equal(t, Map([]int{12, 345, 6789}, istrlen), []int{2, 3, 4})
}

func TestFilter(t *testing.T) {
	assert.Equal(t, Filter([]int{}, isEven), []int{})
	assert.Equal(t, Filter([]int{}, isOdd), []int{})
	assert.Equal(t, Filter([]int{0}, isEven), []int{0})
	assert.Equal(t, Filter([]int{0}, isOdd), []int{})
	assert.Equal(t, Filter([]int{1}, isEven), []int{})
	assert.Equal(t, Filter([]int{1}, isOdd), []int{1})
	assert.Equal(t, Filter([]int{1, 2, 3, 4, 5}, isEven), []int{2, 4})
	assert.Equal(t, Filter([]int{1, 2, 3, 4, 5}, isOdd), []int{1, 3, 5})
}

func TestFilterNot(t *testing.T) {
	assert.Equal(t, FilterNot([]int{}, isEven), []int{})
	assert.Equal(t, FilterNot([]int{}, isOdd), []int{})
	assert.Equal(t, FilterNot([]int{0}, isEven), []int{})
	assert.Equal(t, FilterNot([]int{0}, isOdd), []int{0})
	assert.Equal(t, FilterNot([]int{1}, isEven), []int{1})
	assert.Equal(t, FilterNot([]int{1}, isOdd), []int{})
	assert.Equal(t, FilterNot([]int{1, 2, 3, 4, 5}, isEven), []int{1, 3, 5})
	assert.Equal(t, FilterNot([]int{1, 2, 3, 4, 5}, isOdd), []int{2, 4})
}

func TestForAll(t *testing.T) {
	assert.True(t, ForAll([]int{}, isEven))
	assert.True(t, ForAll([]int{}, isOdd))
	assert.True(t, ForAll([]int{1}, isOdd))
	assert.True(t, ForAll([]int{1, 3}, isOdd))
	assert.True(t, ForAll([]int{2, 4}, isEven))
	assert.False(t, ForAll([]int{1, 2}, isOdd))
	assert.False(t, ForAll([]int{1, 2}, isEven))
}

func TestForAny(t *testing.T) {
	assert.False(t, ForAny([]int{}, isEven))
	assert.False(t, ForAny([]int{}, isOdd))
	assert.True(t, ForAny([]int{1}, isOdd))
	assert.True(t, ForAny([]int{1, 2}, isOdd))
	assert.True(t, ForAny([]int{1, 2}, isEven))
	assert.False(t, ForAny([]int{0, 2, 4}, isOdd))
	assert.False(t, ForAny([]int{1, 3, 5}, isEven))
}

func TestContains(t *testing.T) {
	assert.False(t, Contains([]int{}, 0))
	assert.False(t, Contains([]int{}, 1))
	assert.False(t, Contains([]int{}, -1))
	assert.False(t, Contains([]int{0}, 1))
	assert.False(t, Contains([]int{0}, -1))
	assert.True(t, Contains([]int{0}, 0))
	assert.True(t, Contains([]int{42}, 42))
	assert.True(t, Contains([]int{1, 2, 3}, 2))
}
