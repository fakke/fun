package maybe

import (
	"github.com/stretchr/testify/assert"
	"math"
	"strconv"
	"testing"
)

var itoa = strconv.Itoa
var strlen = func(s string) int { return len(s) }
var izero = func() int { return 0 }
var imul10 = func(a int) int { return a * 10 }
var sempty = func() string { return "" }
var isqrt = func(a int) Maybe[int] {
	if a < 0 {
		return None[int]()
	}
	return Some(int(math.Sqrt(float64(a))))
}

func TestOf(t *testing.T) {
	x1 := 1
	y1 := 1
	x2 := 2
	y2 := 2
	nil1 := (*int)(nil)
	nil2 := (*int)(nil)
	assert.Equal(t, Of(&y1), Of(&x1))
	assert.Equal(t, Of(&y2), Of(&x2))
	assert.Equal(t, Of(nil2), Of(nil1))
	assert.NotEqual(t, Of(&x1), Of(&x2))
	assert.NotEqual(t, Of(&y1), Of(&y2))
	assert.NotEqual(t, Of(&x1), Of(nil1))
}

func TestValue(t *testing.T) {
	x1 := 1
	y1 := 1
	x2 := 2
	y2 := 2
	nil1 := (*int)(nil)
	nil2 := (*int)(nil)
	assert.Equal(t, Some(y1), Some(x1))
	assert.Equal(t, Some(y2), Some(x2))
	assert.Equal(t, Some(nil2), Some(nil1))
	assert.Equal(t, Some(&nil2), Some(&nil1))
	assert.NotEqual(t, Some(&x1), Some(&x2))
	assert.NotEqual(t, Some(&y1), Some(&y2))
	assert.NotEqual(t, Some(&x1), Some(nil1))
}

func TestNothing(t *testing.T) {
	assert.Equal(t, None[int](), None[int]())
	assert.Equal(t, None[string](), None[int]())
}

func TestValue_IsNone(t *testing.T) {
	assert.False(t, Some(42).IsNone())
	assert.False(t, Some("hello").IsNone())
	assert.False(t, Some[any](nil).IsNone())
}

func TestValue_IsSome(t *testing.T) {
	assert.True(t, Some(42).IsSome())
	assert.True(t, Some("hello").IsSome())
	assert.True(t, Some[any](nil).IsSome())
}

func TestValue_Value(t *testing.T) {
	assert.Equal(t, 42, Some(42).Value())
	assert.Equal(t, "hello", Some("hello").Value())
	assert.Equal(t, nil, Some[any](nil).Value())
}

func TestValue_ToNullable(t *testing.T) {
	assert.Equal(t, 42, *Some(42).ToNullable())
	assert.Equal(t, "hello", *Some("hello").ToNullable())
	assert.Equal(t, nil, *Some[any](nil).ToNullable())
}

func TestValue_ValueOr(t *testing.T) {
	assert.Equal(t, 42, Some(42).ValueOr(55))
	assert.Equal(t, "hello", Some("hello").ValueOr("goodbye"))
	assert.Equal(t, nil, Some[any](nil).ValueOr("else"))
}

func TestNothing_IsNone(t *testing.T) {
	assert.True(t, None[int]().IsNone())
	assert.True(t, None[string]().IsNone())
	assert.True(t, None[any]().IsNone())
}

func TestNothing_IsSome(t *testing.T) {
	assert.False(t, None[int]().IsSome())
	assert.False(t, None[string]().IsSome())
	assert.False(t, None[any]().IsSome())
}

func TestNothing_ToNullable(t *testing.T) {
	assert.Nil(t, None[int]().ToNullable())
	assert.Nil(t, None[string]().ToNullable())
	assert.Nil(t, None[any]().ToNullable())
}

func TestNothing_Value(t *testing.T) {
	assert.Panics(t, func() {
		None[int]().Value()
	})
	assert.Panics(t, func() {
		None[string]().Value()
	})
	assert.Panics(t, func() {
		None[any]().Value()
	})
}

func TestNothing_ValueOr(t *testing.T) {
	assert.Equal(t, 42, None[int]().ValueOr(42))
	assert.Equal(t, "hello", None[string]().ValueOr("hello"))
	assert.Equal(t, nil, None[any]().ValueOr(nil))
}

func TestFoldL(t *testing.T) {
	assert.Equal(t, "42", FoldL(Some(42), "", itoa))
	assert.Equal(t, "", FoldL(None[int](), "", itoa))
	assert.Equal(t, 5, FoldL(Some("hello"), 0, strlen))
	assert.Equal(t, 0, FoldL(None[string](), 0, strlen))
}

func TestFoldR(t *testing.T) {
	assert.Equal(t, "42", FoldR(Some(42), sempty, itoa))
	assert.Equal(t, "", FoldR(None[int](), sempty, itoa))
	assert.Equal(t, len("hello"), FoldR(Some("hello"), izero, strlen))
	assert.Equal(t, 0, FoldR(None[string](), izero, strlen))
}

func TestBind(t *testing.T) {
	assert.Equal(t, Some(2), Bind(Some(4), isqrt))
	assert.Equal(t, None[int](), Bind(Some(-4), isqrt))
}

func TestMap(t *testing.T) {
	assert.Equal(t, Some(420), Map(Some(42), imul10))
	assert.Equal(t, Some("42"), Map(Some(42), itoa))
}
