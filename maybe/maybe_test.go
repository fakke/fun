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
var fzero = func() float32 { return 0 }
var imul10 = func(a int) int { return a * 10 }
var fmul3p3 = func(a int) float32 { return float32(a) * 3.3 }
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
	assert.Equal(t, Of(&x1), Of(&y1))
	assert.Equal(t, Of(&x2), Of(&y2))
	assert.Equal(t, Of(nil1), Of(nil2))
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
	assert.Equal(t, Some(x1), Some(y1))
	assert.Equal(t, Some(x2), Some(y2))
	assert.Equal(t, Some(nil1), Some(nil2))
	assert.Equal(t, Some(&nil1), Some(&nil2))
	assert.NotEqual(t, Some(&x1), Some(&x2))
	assert.NotEqual(t, Some(&y1), Some(&y2))
	assert.NotEqual(t, Some(&x1), Some(nil1))
}

func TestNothing(t *testing.T) {
	assert.Equal(t, None[int](), None[int]())
	assert.Equal(t, None[int](), None[string]())
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
	assert.Equal(t, Some(42).Value(), 42)
	assert.Equal(t, Some("hello").Value(), "hello")
	assert.Equal(t, Some[any](nil).Value(), nil)
}

func TestValue_ToNullable(t *testing.T) {
	assert.Equal(t, *Some(42).ToNullable(), 42)
	assert.Equal(t, *Some("hello").ToNullable(), "hello")
	assert.Equal(t, *Some[any](nil).ToNullable(), nil)
}

func TestValue_ValueOr(t *testing.T) {
	assert.Equal(t, Some(42).ValueOr(55), 42)
	assert.Equal(t, Some("hello").ValueOr("goodbye"), "hello")
	assert.Equal(t, Some[any](nil).ValueOr("else"), nil)
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
	assert.Equal(t, None[int]().ValueOr(42), 42)
	assert.Equal(t, None[string]().ValueOr("hello"), "hello")
	assert.Equal(t, None[any]().ValueOr(nil), nil)
}

func TestFoldL(t *testing.T) {
	assert.Equal(t, FoldL(Some(42), "", itoa), "42")
	assert.Equal(t, FoldL(None[int](), "", itoa), "")
	assert.Equal(t, FoldL(Some("hello"), 0, strlen), 5)
	assert.Equal(t, FoldL(None[string](), 0, strlen), 0)
}

func TestFoldR(t *testing.T) {
	assert.Equal(t, FoldR(Some(42), sempty, itoa), "42")
	assert.Equal(t, FoldR(None[int](), sempty, itoa), "")
	assert.Equal(t, FoldR(Some("hello"), izero, strlen), len("hello"))
	assert.Equal(t, FoldR(None[string](), izero, strlen), 0)
}

func TestBind(t *testing.T) {
	assert.Equal(t, Bind(Some(4), isqrt), Some(2))
	assert.Equal(t, Bind(Some(-4), isqrt), None[int]())
}

func TestMap(t *testing.T) {
	assert.Equal(t, Map(Some(42), imul10), Some(420))
	assert.Equal(t, Map(Some(42), itoa), Some("42"))
}
