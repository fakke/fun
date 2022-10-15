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
	assert.Equal(t, -1, Index([]int{}, isOdd))
	assert.Equal(t, -1, Index([]int{0}, isOdd))
	assert.Equal(t, 0, Index([]int{1}, isOdd))
	assert.Equal(t, -1, Index([]int{42}, isOdd))
	assert.Equal(t, -1, Index([]int{2, 4, 6}, isOdd))
	assert.Equal(t, 0, Index([]int{1, 4, 6}, isOdd))
	assert.Equal(t, 1, Index([]int{2, 1, 6}, isOdd))
	assert.Equal(t, 2, Index([]int{2, 4, 1}, isOdd))
}

func TestIndexOf(t *testing.T) {
	assert.Equal(t, -1, IndexOf([]int{}, 0))
	assert.Equal(t, 0, IndexOf([]int{0}, 0))
	assert.Equal(t, -1, IndexOf([]int{1, 2, 3}, 0))
	assert.Equal(t, 0, IndexOf([]int{1, 2, 3}, 1))
	assert.Equal(t, 1, IndexOf([]int{1, 2, 3}, 2))
	assert.Equal(t, 2, IndexOf([]int{1, 2, 3}, 3))
}

func TestFoldL(t *testing.T) {
	// empty
	assert.Equal(t, 0, FoldL([]int{}, 0, iadd))
	assert.Equal(t, 1, FoldL([]int{}, 1, iadd))
	assert.Equal(t, 42, FoldL([]int{}, 42, iadd))
	assert.Equal(t, -1, FoldL([]int{}, -1, iadd))
	assert.Equal(t, -42, FoldL([]int{}, -42, iadd))
	// zero
	assert.Equal(t, 0, FoldL([]int{0}, 0, iadd))
	assert.Equal(t, 1, FoldL([]int{0}, 1, iadd))
	assert.Equal(t, 42, FoldL([]int{0}, 42, iadd))
	assert.Equal(t, -1, FoldL([]int{0}, -1, iadd))
	assert.Equal(t, -42, FoldL([]int{0}, -42, iadd))
	// several zeroes
	assert.Equal(t, 0, FoldL([]int{0, 0}, 0, iadd))
	assert.Equal(t, 1, FoldL([]int{0, 0}, 1, iadd))
	assert.Equal(t, 42, FoldL([]int{0, 0}, 42, iadd))
	assert.Equal(t, -1, FoldL([]int{0, 0}, -1, iadd))
	assert.Equal(t, -42, FoldL([]int{0, 0}, -42, iadd))
	// one value
	assert.Equal(t, 42, FoldL([]int{42}, 0, iadd))
	assert.Equal(t, 43, FoldL([]int{42}, 1, iadd))
	assert.Equal(t, 84, FoldL([]int{42}, 42, iadd))
	assert.Equal(t, 41, FoldL([]int{42}, -1, iadd))
	assert.Equal(t, 0, FoldL([]int{42}, -42, iadd))
	// several values
	assert.Equal(t, 6, FoldL([]int{1, 2, 3}, 0, iadd))
	assert.Equal(t, 7, FoldL([]int{1, 2, 3}, 1, iadd))
	assert.Equal(t, 48, FoldL([]int{1, 2, 3}, 42, iadd))
	assert.Equal(t, 5, FoldL([]int{1, 2, 3}, -1, iadd))
	assert.Equal(t, -36, FoldL([]int{1, 2, 3}, -42, iadd))
	// non-commutative
	assert.Equal(t, []int{}, FoldL([]int{}, []int{}, appendInt))
	assert.Equal(t, []int{0}, FoldL([]int{}, []int{0}, appendInt))
	assert.Equal(t, []int{0}, FoldL([]int{0}, []int{}, appendInt))
	assert.Equal(t, []int{1, 2}, FoldL([]int{2}, []int{1}, appendInt))
	assert.Equal(t, []int{1, 2, 3, 4}, FoldL([]int{3, 4}, []int{1, 2}, appendInt))
}

func TestFoldR(t *testing.T) {
	// empty
	assert.Equal(t, 0, FoldR([]int{}, 0, iadd))
	assert.Equal(t, 1, FoldR([]int{}, 1, iadd))
	assert.Equal(t, 42, FoldR([]int{}, 42, iadd))
	assert.Equal(t, -1, FoldR([]int{}, -1, iadd))
	assert.Equal(t, -42, FoldR([]int{}, -42, iadd))
	// zero
	assert.Equal(t, 0, FoldR([]int{0}, 0, iadd))
	assert.Equal(t, 1, FoldR([]int{0}, 1, iadd))
	assert.Equal(t, 42, FoldR([]int{0}, 42, iadd))
	assert.Equal(t, -1, FoldR([]int{0}, -1, iadd))
	assert.Equal(t, -42, FoldR([]int{0}, -42, iadd))
	// several zeroes
	assert.Equal(t, 0, FoldR([]int{0, 0}, 0, iadd))
	assert.Equal(t, 1, FoldR([]int{0, 0}, 1, iadd))
	assert.Equal(t, 42, FoldR([]int{0, 0}, 42, iadd))
	assert.Equal(t, -1, FoldR([]int{0, 0}, -1, iadd))
	assert.Equal(t, -42, FoldR([]int{0, 0}, -42, iadd))
	// one value
	assert.Equal(t, 42, FoldR([]int{42}, 0, iadd))
	assert.Equal(t, 43, FoldR([]int{42}, 1, iadd))
	assert.Equal(t, 84, FoldR([]int{42}, 42, iadd))
	assert.Equal(t, 41, FoldR([]int{42}, -1, iadd))
	assert.Equal(t, 0, FoldR([]int{42}, -42, iadd))
	// several values
	assert.Equal(t, 6, FoldR([]int{1, 2, 3}, 0, iadd))
	assert.Equal(t, 7, FoldR([]int{1, 2, 3}, 1, iadd))
	assert.Equal(t, 48, FoldR([]int{1, 2, 3}, 42, iadd))
	assert.Equal(t, 5, FoldR([]int{1, 2, 3}, -1, iadd))
	assert.Equal(t, -36, FoldR([]int{1, 2, 3}, -42, iadd))
	// non-commutative
	assert.Equal(t, []int{}, FoldR([]int{}, []int{}, intAppend))
	assert.Equal(t, []int{0}, FoldR([]int{}, []int{0}, intAppend))
	assert.Equal(t, []int{0}, FoldR([]int{0}, []int{}, intAppend))
	assert.Equal(t, []int{1, 2}, FoldR([]int{2}, []int{1}, intAppend))
	assert.Equal(t, []int{1, 2, 4, 3}, FoldR([]int{3, 4}, []int{1, 2}, intAppend))
}

func TestScanL(t *testing.T) {
	// empty
	assert.Equal(t, []int{0}, ScanL([]int{}, 0, iadd))
	assert.Equal(t, []int{1}, ScanL([]int{}, 1, iadd))
	assert.Equal(t, []int{42}, ScanL([]int{}, 42, iadd))
	assert.Equal(t, []int{-1}, ScanL([]int{}, -1, iadd))
	assert.Equal(t, []int{-42}, ScanL([]int{}, -42, iadd))
	// zero
	assert.Equal(t, []int{0, 0}, ScanL([]int{0}, 0, iadd))
	assert.Equal(t, []int{1, 1}, ScanL([]int{0}, 1, iadd))
	assert.Equal(t, []int{42, 42}, ScanL([]int{0}, 42, iadd))
	assert.Equal(t, []int{-1, -1}, ScanL([]int{0}, -1, iadd))
	assert.Equal(t, []int{-42, -42}, ScanL([]int{0}, -42, iadd))
	// several zeroes
	assert.Equal(t, []int{0, 0, 0}, ScanL([]int{0, 0}, 0, iadd))
	assert.Equal(t, []int{1, 1, 1}, ScanL([]int{0, 0}, 1, iadd))
	assert.Equal(t, []int{42, 42, 42}, ScanL([]int{0, 0}, 42, iadd))
	assert.Equal(t, []int{-1, -1, -1}, ScanL([]int{0, 0}, -1, iadd))
	assert.Equal(t, []int{-42, -42, -42}, ScanL([]int{0, 0}, -42, iadd))
	// one value
	assert.Equal(t, []int{0, 42}, ScanL([]int{42}, 0, iadd))
	assert.Equal(t, []int{1, 43}, ScanL([]int{42}, 1, iadd))
	assert.Equal(t, []int{42, 84}, ScanL([]int{42}, 42, iadd))
	assert.Equal(t, []int{-1, 41}, ScanL([]int{42}, -1, iadd))
	assert.Equal(t, []int{-42, 0}, ScanL([]int{42}, -42, iadd))
	// several values
	assert.Equal(t, []int{0, 1, 3, 6}, ScanL([]int{1, 2, 3}, 0, iadd))
	assert.Equal(t, []int{1, 2, 4, 7}, ScanL([]int{1, 2, 3}, 1, iadd))
	assert.Equal(t, []int{42, 43, 45, 48}, ScanL([]int{1, 2, 3}, 42, iadd))
	assert.Equal(t, []int{-1, 0, 2, 5}, ScanL([]int{1, 2, 3}, -1, iadd))
	assert.Equal(t, []int{-42, -41, -39, -36}, ScanL([]int{1, 2, 3}, -42, iadd))
}

func TestScanR(t *testing.T) {
	// empty
	assert.Equal(t, []int{0}, ScanR([]int{}, 0, iadd))
	assert.Equal(t, []int{1}, ScanR([]int{}, 1, iadd))
	assert.Equal(t, []int{42}, ScanR([]int{}, 42, iadd))
	assert.Equal(t, []int{-1}, ScanR([]int{}, -1, iadd))
	assert.Equal(t, []int{-42}, ScanR([]int{}, -42, iadd))
	// zero
	assert.Equal(t, []int{0, 0}, ScanR([]int{0}, 0, iadd))
	assert.Equal(t, []int{1, 1}, ScanR([]int{0}, 1, iadd))
	assert.Equal(t, []int{42, 42}, ScanR([]int{0}, 42, iadd))
	assert.Equal(t, []int{-1, -1}, ScanR([]int{0}, -1, iadd))
	assert.Equal(t, []int{-42, -42}, ScanR([]int{0}, -42, iadd))
	// several zeroes
	assert.Equal(t, []int{0, 0, 0}, ScanR([]int{0, 0}, 0, iadd))
	assert.Equal(t, []int{1, 1, 1}, ScanR([]int{0, 0}, 1, iadd))
	assert.Equal(t, []int{42, 42, 42}, ScanR([]int{0, 0}, 42, iadd))
	assert.Equal(t, []int{-1, -1, -1}, ScanR([]int{0, 0}, -1, iadd))
	assert.Equal(t, []int{-42, -42, -42}, ScanR([]int{0, 0}, -42, iadd))
	// one value
	assert.Equal(t, []int{42, 0}, ScanR([]int{42}, 0, iadd))
	assert.Equal(t, []int{43, 1}, ScanR([]int{42}, 1, iadd))
	assert.Equal(t, []int{84, 42}, ScanR([]int{42}, 42, iadd))
	assert.Equal(t, []int{41, -1}, ScanR([]int{42}, -1, iadd))
	assert.Equal(t, []int{0, -42}, ScanR([]int{42}, -42, iadd))
	// several values
	assert.Equal(t, []int{6, 5, 3, 0}, ScanR([]int{1, 2, 3}, 0, iadd))
	assert.Equal(t, []int{7, 6, 4, 1}, ScanR([]int{1, 2, 3}, 1, iadd))
	assert.Equal(t, []int{48, 47, 45, 42}, ScanR([]int{1, 2, 3}, 42, iadd))
	assert.Equal(t, []int{5, 4, 2, -1}, ScanR([]int{1, 2, 3}, -1, iadd))
	assert.Equal(t, []int{-36, -37, -39, -42}, ScanR([]int{1, 2, 3}, -42, iadd))
}

func TestMap(t *testing.T) {
	assert.Equal(t, []string{}, Map([]int{}, itoa))
	assert.Equal(t, []string{"0"}, Map([]int{0}, itoa))
	assert.Equal(t, []string{"42"}, Map([]int{42}, itoa))
	assert.Equal(t, []string{"1", "2", "3"}, Map([]int{1, 2, 3}, itoa))
	assert.Equal(t, []int{1, 1, 1}, Map([]int{1, 2, 3}, istrlen))
	assert.Equal(t, []int{1, 1, 1}, Map([]int{1, 2, 3}, istrlen))
	assert.Equal(t, []int{2, 3, 4}, Map([]int{12, 345, 6789}, istrlen))
}

func TestFilter(t *testing.T) {
	assert.Equal(t, []int{}, Filter([]int{}, isEven))
	assert.Equal(t, []int{}, Filter([]int{}, isOdd))
	assert.Equal(t, []int{0}, Filter([]int{0}, isEven))
	assert.Equal(t, []int{}, Filter([]int{0}, isOdd))
	assert.Equal(t, []int{}, Filter([]int{1}, isEven))
	assert.Equal(t, []int{1}, Filter([]int{1}, isOdd))
	assert.Equal(t, []int{2, 4}, Filter([]int{1, 2, 3, 4, 5}, isEven))
	assert.Equal(t, []int{1, 3, 5}, Filter([]int{1, 2, 3, 4, 5}, isOdd))
}

func TestFilterNot(t *testing.T) {
	assert.Equal(t, []int{}, FilterNot([]int{}, isEven))
	assert.Equal(t, []int{}, FilterNot([]int{}, isOdd))
	assert.Equal(t, []int{}, FilterNot([]int{0}, isEven))
	assert.Equal(t, []int{0}, FilterNot([]int{0}, isOdd))
	assert.Equal(t, []int{1}, FilterNot([]int{1}, isEven))
	assert.Equal(t, []int{}, FilterNot([]int{1}, isOdd))
	assert.Equal(t, []int{1, 3, 5}, FilterNot([]int{1, 2, 3, 4, 5}, isEven))
	assert.Equal(t, []int{2, 4}, FilterNot([]int{1, 2, 3, 4, 5}, isOdd))
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
